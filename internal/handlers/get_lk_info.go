package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetLKRequest struct {
	// TODO
}

type GetLKResponse struct {
	// TODO
}

func GetLK() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := &GetLKRequest{}
		err := c.Bind(r)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// var lki models.LKInfo
		// _, err = models.GetLKInfo()
		// if err != nil {
		// 	c.String(http.StatusInternalServerError, err.Error())
		// 	return
		// }

		c.JSON(http.StatusOK, GetLKResponse{})
	}
}
