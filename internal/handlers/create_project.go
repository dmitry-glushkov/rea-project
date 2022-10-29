package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateProjectRequest struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Owner string `json:"owner"`
}

type CreateProjectResponse struct{}

// CreateProject ...
func (impl *Implementation) CreateProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &CreateProjectRequest{}
		err := c.BindJSON(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		project := models.Project{
			Name:  r.Name,
			Desc:  r.Desc,
			Owner: r.Owner,
		}
		err = project.Save(c.Request.Context(), impl.DB)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, CreateProjectResponse{})
	}
}
