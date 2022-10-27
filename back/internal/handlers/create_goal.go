package handlers

import (
	"fmt"
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

// CreateGoal ...
func (impl *Implementation) CreateGoal() gin.HandlerFunc {
	type req struct {
		PID     int `json:"pid"`
		Target  int `json:"target"`
		DueDate int `json:"due_date"`
	}

	return func(c *gin.Context) {
		r := &req{}
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

		fmt.Println("HERE")

		fmt.Println(goal)

		err = goal.Save(c.Request.Context(), impl.DB)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}
}
