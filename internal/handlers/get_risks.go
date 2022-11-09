package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type GetRisksRequest struct {
	Pid int `json:"pid"`
}

type GetRisksResponse struct {
	Risks []models.Risk `json:"risks"`
}

func (impl *Implementation) GetRisks() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &GetRisksRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		var (
			risks []models.Risk
		)
		// risks, err = models.GetRisks(c.Request.Context(), impl.DB, r.Pid)
		risks, err = models.GetRisksMock(c.Request.Context(), impl.DB, r.Pid) // TODO mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, GetRisksResponse{
			Risks: risks,
		})
	}
}
