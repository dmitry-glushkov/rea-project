package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

// CreateUser ...
func (impl *Implementation) CreateUser() gin.HandlerFunc {
	type req struct {
		Login string `json:"login"`
		Pass  string `json:"pass"`
		Role  string `json:"role"`
	}

	return func(c *gin.Context) {
		r := &req{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		user := models.User{
			Login: r.Login,
			Pass:  r.Pass,
			Role:  r.Role,
		}
		err = user.Save(c.Request.Context(), impl.DB)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.Status(http.StatusOK)
	}
}
