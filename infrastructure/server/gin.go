package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-4/handler"
	"github.com/hacktiv8-ks07-g04/final-project-4/infrastructure/database"
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
	}

	return router
}

func Run() {
	router := Setup()
	router.Run()
}
