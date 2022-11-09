package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type GetContractorsRequest struct {
	Pid int `json:"pid"`
}

type GetContractorsResponse struct {
	Contractors []models.Contractor `json:"contractors"`
}

func (impl *Implementation) GetContractors() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &GetContractorsRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		var (
			contractors []models.Contractor
		)
		// contractors, err = models.GetContractors(c.Request.Context(), impl.DB, r.Pid)
		contractors, err = models.GetContractorsMock(c.Request.Context(), impl.DB, r.Pid) // TODO mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, GetContractorsResponse{
			Contractors: contractors,
		})
	}
}
