package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type GetInvestorsRequest struct {
	PID int `json:"pid" form:"pid"`
}

type GetInvestorsResponse struct {
	PID       int               `json:"pid"`
	Investors []models.Investor `json:"investors"`
}

func (impl *Implementation) GetInvestors() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &GetInvestorsRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		var investors []models.Investor
		investors, err = models.GetInvestors(c.Request.Context(), impl.DB, r.PID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		resp := GetInvestorsResponse{
			PID:       r.PID,
			Investors: investors,
		}

		c.JSON(http.StatusOK, resp)
	}
}
