package handlers

import (
	"fmt"
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateContractorRequest struct {
	Name     string `json:"name"`
	Interest string `json:"interests"`
}

type CreateContractorResponse struct{}

func (impl *Implementation) CreateContractor() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &CreateContractorRequest{}
		err := c.BindJSON(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		contr := models.Contractor{
			Name:     r.Name,
			Interest: r.Interest,
		}
		fmt.Println(contr)
		err = contr.Save(c.Request.Context(), impl.DB)
		// err = contr.SaveMock(c.Request.Context(), impl.DB) // todo mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, CreateContractorResponse{})
	}
}
