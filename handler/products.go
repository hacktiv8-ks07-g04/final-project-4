package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-4/service"
)

type Products interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type ProductsImpl struct {
	productsService service.Products
}

func InitProducts(service service.Products) *ProductsImpl {
	return &ProductsImpl{service}
}

func (h *ProductsImpl) Create(c *gin.Context) {
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

func (h *ProductsImpl) GetAll(c *gin.Context) {
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

func (h *ProductsImpl) Update(c *gin.Context) {
	body := dto.UpdateProductRequest{}
	id := c.Param("productId")

	log.Print("id ", id)
	log.Print("body ", body)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errs.BadRequest("invalid request body"))
		return
	}

	product, err := h.productsService.Update(id, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errs.InternalServerError(err.Error()))
		return
	}

	response := dto.UpdateProductResponse{
		ID:         product.ID,
		Title:      product.Title,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (h *ProductsImpl) Delete(c *gin.Context) {
	id := c.Param("productId")

	err := h.productsService.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, errs.NotFound("product not found"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product has been successfully deleted",
	})
}
