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
