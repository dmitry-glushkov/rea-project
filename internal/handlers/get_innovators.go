package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type GetInnovatorsRequest struct{}

type GetInnovatorsResponse struct {
	Innovators []models.Innovator `json:"innovators"`
}

func (impl *Implementation) GetInnovators() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &GetInnovatorsRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		var (
			innovators []models.Innovator
		)
		// innovators, err = models.GetInnovators(c.Request.Context(), impl.DB)
		innovators, err = models.GetInnovatorsMock(c.Request.Context(), impl.DB) // TODO mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, GetInnovatorsResponse{
			Innovators: innovators,
		})
	}
}
