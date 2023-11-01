package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-4/handler"
	"github.com/hacktiv8-ks07-g04/final-project-4/infrastructure/database"
	"github.com/hacktiv8-ks07-g04/final-project-4/middleware"
	"github.com/hacktiv8-ks07-g04/final-project-4/repository"
	"github.com/hacktiv8-ks07-g04/final-project-4/service"
)

func Setup() *gin.Engine {
	router := gin.Default()
	db := database.GetInstance()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "active",
			"message": "Welcome to the Toko Belanja API",
			"version": "1.0.0",
		})
	})

	// Users
	usersRepo := repository.UsersRepositoryInit(db)
	usersService := service.UsersServiceInit(usersRepo)
	usersHandler := handler.UsersHandlerInit(usersService)

	users := router.Group("/users")
	{
		users.POST("/register", usersHandler.Register)
		users.POST("/login", usersHandler.Login)
		users.PATCH("/topup", middleware.Authentication(), usersHandler.TopUp)
	}

	// Categories
	categoriesRepo := repository.CategoriesRepositoryInit(db)
	categoriesService := service.CategoriesServiceInit(categoriesRepo)
	categoriesHandler := handler.CategoriesHandlerInit(categoriesService)

	categories := router.Group("/categories").Use(middleware.Authentication())
	{
		categories.POST("/", middleware.AdminAuthorization(), categoriesHandler.Create)
		categories.GET("/", categoriesHandler.GetAll)
		categories.PATCH("/:categoryId", middleware.AdminAuthorization(), categoriesHandler.Update)
		categories.DELETE("/:categoryId", middleware.AdminAuthorization(), categoriesHandler.Delete)
	}

	// Products
	productsRepo := repository.ProductsRepositoryInit(db)
	productsService := service.ProductsServiceInit(productsRepo)
	productsHandler := handler.ProductsHandlerInit(productsService)

	products := router.Group("/products").Use(middleware.Authentication())
	{
		products.POST("/", middleware.AdminAuthorization(), productsHandler.Create)
		products.GET("/", productsHandler.GetAll)
		products.PATCH("/:productId", productsHandler.Update)
	}

	// Auth Purpose
	router.GET("/auth", middleware.Authentication(), func(c *gin.Context) {
		user := c.MustGet("user").(map[string]interface{})

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})

	return router
}

func Run() {
	router := Setup()
	router.Run()
}
