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
				Owner:       "aboba",
				ProjectName: "biba",
				Sum:         178,
				Target:      300,
			},
			{
				PID:         2,
				Owner:       "aboba vtoraya",
				ProjectName: "biba vtoray",
				Sum:         1780,
				Target:      3000,
			},
			{
				PID:         3,
				Owner:       "aboba tri",
				ProjectName: "biba tri",
				Sum:         17800,
				Target:      30000,
			},
			{
				PID:         4,
				Owner:       "aboba chi da",
				ProjectName: "biba chi da",
				Sum:         178000,
				Target:      300000,
			},
			{
				PID:         5,
				Owner:       "aboba pyat'",
				ProjectName: "biba pyat'",
				Sum:         1780000,
				Target:      3000000,
			},
		},
	}
}
