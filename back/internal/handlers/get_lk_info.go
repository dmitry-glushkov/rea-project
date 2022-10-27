package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLKInfo() gin.HandlerFunc {
	type req struct {
	}

	return func(c *gin.Context) {
		r := &req{}
		err := c.BindJSON(r)
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
	}
}
