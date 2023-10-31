package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-4/service"
)

type UsersHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	TopUp(c *gin.Context)
}

type UsersHandlerImpl struct {
	usersService service.UsersService
}

func UsersHandlerInit(service service.UsersService) *UsersHandlerImpl {
	return &UsersHandlerImpl{service}
}

func (u *UsersHandlerImpl) Register(c *gin.Context) {
	body := dto.RegisterRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errs.BadRequest("Invalid request body"))
		return
	}

	user, err := u.usersService.Register(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errs.InternalServerError(err.Error()))
		return
	}

	response := dto.RegisterResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  user.Password,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (u *UsersHandlerImpl) Login(c *gin.Context) {
	body := dto.LoginRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errs.BadRequest("Invalid request body"))
		return
	}

	token, err := u.usersService.Login(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errs.InternalServerError(err.Error()))
		return
	}

	response := dto.LoginResponse{
		Token: token,
	}

	c.JSON(http.StatusOK, response)
}

func (u *UsersHandlerImpl) TopUp(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	body := dto.TopUpRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errs.BadRequest("Invalid request body"))
		return
	}

	balance, err := u.usersService.TopUp(user["id"].(uint), body.Balance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errs.InternalServerError(err.Error()))
		return
	}

	response := dto.TopUpResponse{
		Message: fmt.Sprintf("Your balance has been successfully updated to Rp %d", balance),
	}

	c.JSON(http.StatusOK, response)
}
