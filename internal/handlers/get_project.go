package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type GetProjectRequest struct {
	PID int `json:"pid" form:"pid"`
}

type GetProjectResponse struct {
	Project models.Project `json:"project_info"`
	Stages  []models.Stage `json:"project_stages"`
}

func (impl *Implementation) GetProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &GetProjectRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		var project models.Project
		project, err = models.GetProjectMock(c.Request.Context(), impl.DB, r.PID) // TODO mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		var stages []models.Stage
		stages, err = models.GetStagesMock(c.Request.Context(), impl.DB, r.PID) // TODO mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, GetProjectResponse{
			Project: project,
			Stages:  stages,
		})
	}
}
