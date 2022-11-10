package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateInvestorRequest struct {
	Name      string `json:"name"`
	Interests string `json:"interests"`
}

type CreateInvestorResponse struct{}

// CreateInvestor ...
func (impl *Implementation) CreateInvestor() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &CreateInvestorRequest{}
		err := c.BindJSON(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		investor := models.Investor{
			Name:      r.Name,
			Interests: r.Interests,
		}
		err = investor.Save(c.Request.Context(), impl.DB)
		// err = investor.SaveMock(c.Request.Context(), impl.DB) // todo mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, CreateInvestorResponse{})
	}
}
