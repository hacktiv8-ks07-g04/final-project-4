package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-4/service"
)

type ProductsHandler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
}

type ProductsHandlerImpl struct {
	productsService service.ProductsService
}

func ProductsHandlerInit(service service.ProductsService) *ProductsHandlerImpl {
	return &ProductsHandlerImpl{service}
}

func (h *ProductsHandlerImpl) Create(c *gin.Context) {
	body := dto.CreateProductRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errs.BadRequest("invalid request body"))
		return
	}

	product, err := h.productsService.Create(body)
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

func (h *ProductsHandlerImpl) GetAll(c *gin.Context) {
	products, err := h.productsService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errs.InternalServerError(err.Error()))
		return
	}

	response := []dto.CreateProductResponse{}

	for _, product := range products {
		response = append(response, dto.CreateProductResponse{
			ID:         product.ID,
			Title:      product.Title,
			Price:      product.Price,
			Stock:      product.Stock,
			CategoryID: product.CategoryID,
			CreatedAt:  product.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}
