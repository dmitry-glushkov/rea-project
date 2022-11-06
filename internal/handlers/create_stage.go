package handlers

import (
	"net/http"
	"time"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateStageRequest struct {
	PID     int    `json:"pid"`
	Target  int    `json:"target"`
	DueDate string `json:"due_date"`
}

type CreateStageResponse struct{}

// CreateStage ...
func (impl *Implementation) CreateStage() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &CreateStageRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		dd, err := time.Parse("02.01.2006", r.DueDate)
		if err != nil {
			// TODO
			return
		}

		goal := models.Stage{
			PID:     r.PID,
			Target:  r.Target,
			DueDate: dd.Format("02.01.2006"),
		}
		err = goal.Save(c.Request.Context(), impl.DB)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, CreateStageResponse{})
	}
}
