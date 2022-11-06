package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateInvestementRequest struct {
	Uid int `json:"uid"`
	Pid int `json:"pid"`
	Val int `json:"val"`
}

type CreateInvestementResponse struct{}

func (impl *Implementation) CreateInvestement() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &CreateInvestementRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		investement := models.Investement{
			UID: r.Uid,
			PID: r.Pid,
			Val: r.Val,
		}
		err = investement.Save(c.Request.Context(), impl.DB)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, CreateInvestementResponse{})
	}
}
