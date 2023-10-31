package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hacktiv8-ks07-g04/final-project-4/pkg/errs"
)

// check if the user role is admin
func AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(map[string]interface{})
		role := user["role"].(string)

		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errs.Unauthorized("Unauthorized"))
		}

		c.Next()
	}
}
