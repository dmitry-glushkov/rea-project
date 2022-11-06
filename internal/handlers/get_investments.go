package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type GetInvestmentsRequest struct {
	PID int `json:"pid" form:"pid"`
}

type GetInvestmentsResponse struct {
	Investments []models.Investment `json:"investments"`
}

func (impl *Implementation) GetInvestments() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &GetInvestmentsRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		var investments []models.Investment
		investments, err = models.GetInvestmentsMock(c.Request.Context(), impl.DB, r.PID) // TODO mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		resp := GetInvestmentsResponse{
			Investments: investments,
		}

		c.JSON(http.StatusOK, resp)
	}
}
