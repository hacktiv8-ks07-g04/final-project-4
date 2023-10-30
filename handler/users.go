package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-4/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-4/service"
)

type UsersHandler interface {
	Register(c *gin.Context)
}

type UsersHandlerImpl struct {
	usersService service.UsersService
}

func UsersHandlerInit(service service.UsersService) *UsersHandlerImpl {
	return &UsersHandlerImpl{service}
}

func (u *UsersHandlerImpl) Register(c *gin.Context) {
	user := entity.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errs.BadRequest("Invalid request body"))
		return
	}

	response, err := u.usersService.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errs.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}
