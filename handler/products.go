package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-4/service"
)

type ProductsHandler interface {
	Add(c *gin.Context)
}

type ProductsHandlerImpl struct {
	productsService service.ProductsService
}

func ProductsHandlerInit(service service.ProductsService) *ProductsHandlerImpl {
	return &ProductsHandlerImpl{service}
}

func (h *ProductsHandlerImpl) Add(c *gin.Context) {
	body := dto.CreateProductRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errs.BadRequest("invalid request body"))
		return
	}

	product, err := h.productsService.Add(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errs.InternalServerError(err.Error()))
		return
	}

	response := dto.CreateProductResponse{
		ID:         product.ID,
		Title:      product.Title,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
		CreatedAt:  product.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}
