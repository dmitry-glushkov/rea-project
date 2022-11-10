package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type GetProjectsRequest struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

type GetProjectsResponse struct {
	Projects []models.Project `json:"projects"`
}

func (impl *Implementation) GetProjects() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &GetProjectsRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		var (
			projects []models.Project
		)
		projects, err = models.GetProjects(c.Request.Context(), impl.DB, r.Page, r.Limit)
		// projects, err = models.GetProjectsMock(c.Request.Context(), impl.DB, r.Page, r.Limit) // TODO mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, GetProjectsResponse{
			Projects: projects,
		})
	}
}
