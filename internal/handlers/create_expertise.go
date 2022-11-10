package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateExperiseRequest struct {
	Pid     int    `json:"pid"`
	Content string `json:"content"`
}

type CreateExpertiseResponse struct{}

func (impl *Implementation) CreateExpertise() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &CreateExperiseRequest{}
		err := c.BindJSON(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		exp := models.Expertise{
			Pid:     r.Pid,
			Content: r.Content,
		}
		err = exp.Save(c.Request.Context(), impl.DB)
		// err = exp.SaveMock(c.Request.Context(), impl.DB) // todo mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, CreateExpertiseResponse{})
	}
}
