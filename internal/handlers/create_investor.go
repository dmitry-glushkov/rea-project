package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateInvestorRequest struct {
	Login string `json:"login"`
}

type CreateInvestorResponse struct{}

// CreateInvestor ...
func (impl *Implementation) CreateInvestor() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &CreateInvestorRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		investor := models.Investor{
			Login: r.Login,
		}
		err = investor.Save(c.Request.Context(), impl.DB)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, CreateInvestorResponse{})
	}
}
