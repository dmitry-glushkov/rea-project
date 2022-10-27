package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

func (impl *Implementation) GetProjects() gin.HandlerFunc {
	type req struct {
		Page  int `json:"page" form:"page"`
		Limit int `json:"limit" form:"limit"`
	}

	return func(c *gin.Context) {
		r := &req{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		var projects []models.Project
		projects, err = models.GetProjects(c.Request.Context(), impl.DB, r.Page, r.Limit)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, projects)
	}
}
