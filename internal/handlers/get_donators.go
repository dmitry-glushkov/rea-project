package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type GetDonatorsRequest struct {
	PID int `json:"pid" form:"pid"`
}

type GetDonatorsResponse struct {
	PID int `json:"pid"`
	// TODO
}

func (impl *Implementation) GetDonators() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &GetDonatorsRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// var donators []models.User
		_, err = models.GetDonators(c.Request.Context(), impl.DB, r.PID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		resp := GetDonatorsResponse{}

		c.JSON(http.StatusOK, resp)
	}
}
