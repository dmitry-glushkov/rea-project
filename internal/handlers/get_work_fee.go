package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type GetWorkFeeRequest struct {
	Pid int `json:"pid"`
	Cid int `json:"cid"`
}

type GetWorkFeeResponse struct {
	WFs []models.WorkFee `json:"wfs"`
}

func (impl *Implementation) GetWorkFee() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &GetWorkFeeRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		var (
			wfs []models.WorkFee
		)
		wfs, err = models.GetWFs(c.Request.Context(), impl.DB, r.Pid, r.Cid)
		// wfs, err = models.GetWFsMock(c.Request.Context(), impl.DB, r.Pid, r.Cid) // TODO mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, GetWorkFeeResponse{
			WFs: wfs,
		})
	}
}
