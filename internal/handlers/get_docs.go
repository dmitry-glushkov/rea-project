package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type GetDocsRequest struct {
	Pid int `json:"pid"`
}

type GetDocsResponse struct {
	Docs []models.Doc `json:"docs"`
}

func (impl *Implementation) GetDocs() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &GetDocsRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		var (
			docs []models.Doc
		)
		docs, err = models.GetDocs(c.Request.Context(), impl.DB, r.Pid)
		// docs, err = models.GetDocsMock(c.Request.Context(), impl.DB, r.Pid) // TODO mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, GetDocsResponse{
			Docs: docs,
		})
	}
}
