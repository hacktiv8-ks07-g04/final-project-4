package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-4/service"
)

type Transactions interface {
	Create(c *gin.Context)
	GetUserTransactions(c *gin.Context)
	GetAll(c *gin.Context)
}

type TransactionsImpl struct {
	service service.Transactions
}

func InitTransactions(service service.Transactions) *TransactionsImpl {
	return &TransactionsImpl{service}
}

func (h *TransactionsImpl) Create(c *gin.Context) {
	body := dto.CreateTransactionRequest{}
	user := c.MustGet("user").(map[string]interface{})
	userID := user["id"].(uint)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errs.BadRequest("invalid request body"))
		return
	}

	response, err := h.service.Create(userID, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errs.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *TransactionsImpl) GetUserTransactions(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	userID := user["id"].(uint)

	response, err := h.service.GetUserTransactions(userID)
	if err != nil {
		if err.Error() == "transactions not found" {
			c.JSON(http.StatusNotFound, errs.NotFound("transactions not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, errs.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *TransactionsImpl) GetAll(c *gin.Context) {
	response, err := h.service.GetAll()
	if err != nil {
		if err.Error() == "transactions are empty" {
			c.JSON(http.StatusNotFound, errs.NotFound("transactions are empty"))
			return
		}
		c.JSON(http.StatusInternalServerError, errs.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}
