package handlers

import (
	"net/http"

	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

func (impl *Implementation) GetGoals() gin.HandlerFunc {
	type req struct {
		PID int `json:"pid" form:"pid"`
	}

	return func(c *gin.Context) {
		r := &req{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		var goals []models.Goal
		goals, err = models.GetGoals(c.Request.Context(), impl.DB, r.PID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, goals)
	}
}
