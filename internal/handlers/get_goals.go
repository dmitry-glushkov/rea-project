package handlers

import (
	"net/http"

	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type GetGoalsRequest struct {
	PID int `json:"pid" form:"pid"`
}

type GetGoalsResponse struct {
	PID int `json:"pid" form:"pid"`
	// TODO
}

func (impl *Implementation) GetGoals() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &GetGoalsRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// var goals []models.Goal
		_, err = models.GetGoals(c.Request.Context(), impl.DB, r.PID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		resp := GetGoalsResponse{}

		c.JSON(http.StatusOK, resp)
	}
}
