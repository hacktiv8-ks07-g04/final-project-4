package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-4/service"
)

type CategoriesHandler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
}

type CategoriesHandlerImpl struct {
	categoriesService service.CategoriesService
}

func CategoriesHandlerInit(service service.CategoriesService) *CategoriesHandlerImpl {
	return &CategoriesHandlerImpl{service}
}

func (h *CategoriesHandlerImpl) Create(c *gin.Context) {
	body := dto.CreateCategoryRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errs.BadRequest("invalid request body"))
		return
	}

	category, err := h.categoriesService.Create(body)
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

func (h *CategoriesHandlerImpl) GetAll(c *gin.Context) {
	categories, err := h.categoriesService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errs.InternalServerError(err.Error()))
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
