package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetProjectRequest struct {
	PID int `json:"pid" form:"pid"`
}

type GetProjectResponse struct {
	Project ProjectInfo `json:"project_info"`
}

func (impl *Implementation) GetProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &GetProjectRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// var project models.Project
		// project, err = models.GetProject(c.Request.Context(), impl.DB, r.PID)
		// if err != nil {
		// 	c.String(http.StatusInternalServerError, err.Error())
		// 	return
		// }

		project := createMockProject(r.PID)

		c.JSON(http.StatusOK, GetProjectResponse{
			Project: project,
		})
	}
}

func createMockProject(pid int) ProjectInfo {
	return ProjectInfo{
		PID:   pid,
		Owner: "создатель проекта",
		Desc: `
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		`,
		ProjectName: "название проекта",
		Sum:         100000,
		Target:      400000,
	}
}
