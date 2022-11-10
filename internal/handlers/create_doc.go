package handlers

import (
	"net/http"
	"ucheba/back/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateDocRequest struct {
	Pid    int    `json:"pid"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Doc    string `json:"doc"`
	Cid    string `json:"cid"`
}

type CreateDocResponse struct{}

func (impl *Implementation) CreateDoc() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &CreateDocRequest{}
		err := c.BindJSON(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		doc := models.Doc{
			Pid:   r.Pid,
			Title: r.Title,
			Dcm:   r.Doc,
			Cid:   r.Cid,
		}
		err = doc.Save(c.Request.Context(), impl.DB)
		// err = doc.SaveMock(c.Request.Context(), impl.DB) // todo mock
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, CreateDocResponse{})
	}
}
