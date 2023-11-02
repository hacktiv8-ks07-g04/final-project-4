package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-4/service"
)

type Categories interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type CategoriesImpl struct {
	service service.Categories
}

func InitCategories(service service.Categories) *CategoriesImpl {
	return &CategoriesImpl{service}
}

func (h *CategoriesImpl) Create(c *gin.Context) {
	body := dto.CreateCategoryRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errs.BadRequest("invalid request body"))
		return
	}

	category, err := h.service.Create(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errs.InternalServerError(err.Error()))
		return
	}

	response := dto.CreateCategoryResponse{
		ID:                category.ID,
		Type:              category.Type,
		SoldProductAmount: category.SoldProductAmount,
		CreatedAt:         category.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *CategoriesImpl) GetAll(c *gin.Context) {
	categories, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusNotFound, errs.NotFound("categories not found"))
		return
	}

	var response []dto.GetCategoryResponse
	for _, category := range categories {
		var products []dto.GetProductResponse
		for _, product := range category.Products {
			products = append(products, dto.GetProductResponse{
				ID:        product.ID,
				Title:     product.Title,
				Price:     product.Price,
				Stock:     product.Stock,
				CreatedAt: product.CreatedAt,
				UpdatedAt: product.UpdatedAt,
			})
		}

		response = append(response, dto.GetCategoryResponse{
			ID:                category.ID,
			Type:              category.Type,
			SoldProductAmount: category.SoldProductAmount,
			CreatedAt:         category.CreatedAt,
			UpdatedAt:         category.UpdatedAt,
			Products:          products,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *CategoriesImpl) Update(c *gin.Context) {
	body := dto.CreateCategoryRequest{}
	id := c.Param("categoryId")

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errs.BadRequest("invalid request body"))
		return
	}

	category, err := h.service.Update(id, body.Type)
	if err != nil {
		c.JSON(http.StatusNotFound, errs.NotFound("category not found"))
		return
	}

	response := dto.CreateCategoryResponse{
		ID:                category.ID,
		Type:              category.Type,
		SoldProductAmount: category.SoldProductAmount,
		CreatedAt:         category.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (h *CategoriesImpl) Delete(c *gin.Context) {
	id := c.Param("categoryId")

	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, errs.NotFound("category not found"))
		return
	}

	response := dto.DeleteCategoryResponse{
		Message: "Category has been successfully deleted",
	}

	c.JSON(http.StatusOK, response)
}
