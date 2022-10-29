package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateDonateRequest struct {
	Uid int `json:"uid"`
	Pid int `json:"pid"`
	Val int `json:"val"`
}

type CreateDonateResponse struct{}

func (impl *Implementation) CreateDonate() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &CreateDonateRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		donate := models.Donate{
			UID: r.Uid,
			PID: r.Pid,
			Val: r.Val,
		}
		err = donate.Save(c.Request.Context(), impl.DB)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, CreateDonateResponse{})
	}
}
