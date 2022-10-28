package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

// CreateProject ...
func (impl *Implementation) CreateProject() gin.HandlerFunc {
	type req struct {
		Name    string `json:"name"`
		Desc    string `json:"desc"`
		OwnerID int    `json:"owner_id"`
	}

	return func(c *gin.Context) {
		r := &req{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		project := models.Project{
			Name:    r.Name,
			Desc:    r.Desc,
			OwnerID: r.OwnerID,
		}
		err = project.Save(c.Request.Context(), impl.DB)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.Status(http.StatusOK)
	}
}
