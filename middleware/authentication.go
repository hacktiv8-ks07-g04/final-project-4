package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-4/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-4/utils"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := utils.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errs.Unauthorized(err.Error()))
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
