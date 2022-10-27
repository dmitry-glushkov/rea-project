package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type Implementation struct {
	DB     *pgx.Conn
	Router *gin.Engine
}

func NewImplementation(db *pgx.Conn) Implementation {
	impl := Implementation{
		DB:     db,
		Router: gin.Default(),
	}
	impl.registerHTTP()

	return impl
}

func (impl *Implementation) registerHTTP() {
	// GET
	impl.Router.GET("/api/donators", impl.GetDonators())
	impl.Router.GET("/api/goals", impl.GetGoals())
	impl.Router.GET("/api/project", impl.GetProject())
	impl.Router.GET("/api/projects", impl.GetProjects())

	// CREATE
	impl.Router.POST("/api/create_donate", impl.CreateDonate())
	impl.Router.POST("/api/create_project", impl.CreateProject())
	impl.Router.POST("/api/create_user", impl.CreateUser())
	impl.Router.POST("/api/create_goal", impl.CreateGoal())

	// WEB MORDA
	impl.Router.LoadHTMLGlob("pages/*")
	impl.Router.GET("/projects", func(c *gin.Context) {
		c.HTML(http.StatusOK, "projects.html", gin.H{})
	})
	impl.Router.GET("/project", func(c *gin.Context) {
		c.HTML(http.StatusOK, "project.html", gin.H{})
	})
	impl.Router.GET("/donators", func(c *gin.Context) {
		c.HTML(http.StatusOK, "donators.html", gin.H{})
	})
	impl.Router.GET("/lk", func(c *gin.Context) {
		c.HTML(http.StatusOK, "lk.html", gin.H{})
	})
	impl.Router.GET("/donates", func(c *gin.Context) {
		c.HTML(http.StatusOK, "donates.html", gin.H{})
	})
	impl.Router.GET("/goals", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goals.html", gin.H{})
	})
	impl.Router.GET("/crm", func(c *gin.Context) {
		c.HTML(http.StatusOK, "crm.html", gin.H{})
	})
}

func (impl *Implementation) Run() error {
	return impl.Router.Run("localhost:8080")
}
