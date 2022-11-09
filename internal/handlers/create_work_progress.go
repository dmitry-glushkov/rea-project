package handlers

import "github.com/gin-gonic/gin"

// todo проверить необходимость

type CreateWorkProgressRequest struct {
	Name     string `json:"name"`
	Activity string `json:"activity"`
}

type CreateWorkProgressResponse struct{}

func (impl *Implementation) CreateWorkProgress() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}
