package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateInvestmentRequest struct {
	Uid int `json:"uid"`
	Pid int `json:"pid"`
	Val int `json:"val"`
}

type CreateInvestmentResponse struct{}

func (impl *Implementation) CreateInvestment() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &CreateInvestmentRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		investment := models.Investment{
			UID: r.Uid,
			PID: r.Pid,
			Val: r.Val,
		}
		// err = investment.Save(c.Request.Context(), impl.DB)
		err = investment.SaveMock(c.Request.Context(), impl.DB) // todo mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, CreateInvestmentResponse{})
	}
}
