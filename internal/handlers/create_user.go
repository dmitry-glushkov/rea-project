package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
	Role  string `json:"role"`
}

type CreateUserResponse struct{}

// CreateUser ...
func (impl *Implementation) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &CreateUserRequest{}
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

		c.JSON(http.StatusOK, CreateUserResponse{})
	}
}
