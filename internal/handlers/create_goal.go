package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateGoalRequest struct {
	PID     int `json:"pid"`
	Target  int `json:"target"`
	DueDate int `json:"due_date"`
}

type CreateGoalResponse struct{}

// CreateGoal ...
func (impl *Implementation) CreateGoal() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &CreateGoalRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		goal := models.Goal{
			PID:     r.PID,
			Target:  r.Target,
			DueDate: r.DueDate,
		}
		err = goal.Save(c.Request.Context(), impl.DB)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, CreateGoalResponse{})
	}
}
