package handlers

import (
	"net/http"
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
		err := c.BindJSON(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		goal := models.Stage{
			PID:     r.PID,
			Target:  r.Target,
			DueDate: r.DueDate,
		}
		err = goal.Save(c.Request.Context(), impl.DB)
		// err = goal.SaveMock(c.Request.Context(), impl.DB) // todo mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, CreateStageResponse{})
	}
}
