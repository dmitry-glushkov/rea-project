package handlers

import (
	"net/http"

	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type GetStagesRequest struct {
	PID int `json:"pid" form:"pid"`
}

type GetStagesResponse struct {
	PID    int `json:"pid" form:"pid"`
	Stages []models.Stage
}

func (impl *Implementation) GetStages() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &GetStagesRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		var stages []models.Stage
		stages, err = models.GetStages(c.Request.Context(), impl.DB, r.PID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		resp := GetStagesResponse{
			PID:    r.PID,
			Stages: stages,
		}

		c.JSON(http.StatusOK, resp)
	}
}
