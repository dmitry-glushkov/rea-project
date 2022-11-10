package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateWorkFeeRequest struct {
	Pid        int `json:"pid"`
	Cid        int `json:"cid"`
	Sum        int `json:"sum"`
	DocumentId int `json:"document_id"`
}

type CreateWorkFeeResponse struct{}

func (impl *Implementation) CreateWorkFee() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &CreateWorkFeeRequest{}
		err := c.BindJSON(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		wf := models.WorkFee{
			Pid: r.Pid,
			Cid: r.Cid,
			Sum: r.Sum,
			Dcm: models.Doc{ID: r.DocumentId},
		}
		err = wf.Save(c.Request.Context(), impl.DB)
		// err = wf.SaveMock(c.Request.Context(), impl.DB) // todo mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, CreateWorkFeeResponse{})
	}
}
