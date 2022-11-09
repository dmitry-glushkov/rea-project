package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type GetExpertisesRequest struct {
	Pid int `json:"pid"`
}

type GetExpertisesResponse struct {
	Expertises []models.Expertise `json:"expertises"`
}

func (impl *Implementation) GetExpertise() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &GetExpertisesRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		var (
			expertises []models.Expertise
		)
		// expertises, err = models.GetExpertises(c.Request.Context(), impl.DB, r.Pid)
		expertises, err = models.GetExpertisesMock(c.Request.Context(), impl.DB, r.Pid) // TODO mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, GetExpertisesResponse{
			Expertises: expertises,
		})
	}
}
