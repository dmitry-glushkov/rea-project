package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateInnovatorRequest struct {
	Name string `json:"name"`
}

type CreateInnovatorResponse struct{}

func (impl *Implementation) CreateInnovator() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &CreateInnovatorRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		in := models.Innovator{
			Name: r.Name,
		}
		// err = in.SaveMock(c.Request.Context(), impl.DB)
		err = in.SaveMock(c.Request.Context(), impl.DB) // todo mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, CreateInnovatorResponse{})
	}
}
