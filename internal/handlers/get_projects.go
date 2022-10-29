package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetProjectsRequest struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

type ProjectInfo struct {
	PID         int    `json:"pid"`
	Owner       string `json:"owner"`
	Desc        string `json:"desc"`
	ProjectName string `json:"project_name"`
	Sum         int    `json:"sum"`
	Target      int    `json:"target"`
}

type GetProjectsResponse struct {
	Projects []ProjectInfo `json:"projects_info"`
}

func (impl *Implementation) GetProjects() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &GetProjectsRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// var projects []models.Project
		// projects, err = models.GetProjects(c.Request.Context(), impl.DB, r.Page, r.Limit)
		// if err != nil {
		// 	c.String(http.StatusInternalServerError, err.Error())
		// 	return
		// }

		mockProjects := createMockProjects()

		c.JSON(http.StatusOK, mockProjects)
	}
}

func createMockProjects() GetProjectsResponse {
	return GetProjectsResponse{
		[]ProjectInfo{
			{
				PID:         1,
				Owner:       "владелец 1",
				ProjectName: "проект 1",
				Desc:        "описание проекта 1",
				Sum:         178,
				Target:      300,
			},
			{
				PID:         2,
				Owner:       "владелец 2",
				ProjectName: "проект 2",
				Desc:        "описание описание описание описание описание описание описание описание описание описание описание описание описание описание описание",
				Sum:         1780,
				Target:      3000,
			},
			{
				PID:         3,
				Owner:       "владелец 3",
				ProjectName: "проект 3",
				Sum:         17800,
				Target:      30000,
			},
			{
				PID:         4,
				Owner:       "владелец 4",
				ProjectName: "проект 4",
				Sum:         178000,
				Target:      300000,
			},
			{
				PID:         5,
				Owner:       "владелец 5",
				ProjectName: "проект 5",
				Sum:         1780000,
				Target:      3000000,
			},
		},
	}
}
