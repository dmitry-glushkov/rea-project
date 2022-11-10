package handlers

import (
	"fmt"
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateRiskRequest struct {
	Pid  int    `json:"pid"`
	Risk string `json:"risk"`
	Plan string `json:"plan"`
	Sum  int    `json:"sum"`
}

type CreateRiskResponse struct{}

func (impl *Implementation) CreateRisk() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &CreateRiskRequest{}
		err := c.BindJSON(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		risk := models.Risk{
			PID:  r.Pid,
			Rsk:  r.Risk,
			Plan: r.Plan,
			Sum:  r.Sum,
		}
		fmt.Println(risk)
		err = risk.Save(c.Request.Context(), impl.DB)
		// err = risk.SaveMock(c.Request.Context(), impl.DB) // todo mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, CreateRiskResponse{})
	}
}
