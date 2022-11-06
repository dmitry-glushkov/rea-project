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
	impl.Router.GET("/api/stages", impl.GetStages())
	impl.Router.GET("/api/project", impl.GetProject())
	impl.Router.GET("/api/projects", impl.GetProjects())
	impl.Router.GET("/api/investments", impl.GetInvestments())
	impl.Router.GET("/api/investors", impl.GetInvestors())
	impl.Router.GET("/api/innovators", impl.GetInnovators())
	impl.Router.GET("/api/contractors", impl.GetContractors())
	impl.Router.GET("/api/work_progress", impl.GetWorkProgress())
	impl.Router.GET("/api/docs", impl.GetDocs())
	impl.Router.GET("/api/risks", impl.GetRisks())

	// CREATE
	impl.Router.POST("/api/create_investment", impl.CreateInvestment())
	impl.Router.POST("/api/create_project", impl.CreateProject())
	impl.Router.POST("/api/create_investor", impl.CreateInvestor())
	impl.Router.POST("/api/create_innovator", impl.CreateInnovator())
	impl.Router.POST("/api/create_contractor", impl.CreateContractor())
	impl.Router.POST("/api/create_stage", impl.CreateStage())
	impl.Router.POST("/api/create_work_progress", impl.CreateWorkProgress())
	impl.Router.POST("/api/create_doc", impl.CreateDoc())
	impl.Router.POST("/api/create_risk", impl.CreateRisk())

	// WEB MORDA
	impl.Router.LoadHTMLGlob("pages/*")
	impl.Router.GET("/projects", func(c *gin.Context) {
		c.HTML(http.StatusOK, "projects.html", gin.H{})
	})
	impl.Router.GET("/project", func(c *gin.Context) {
		c.HTML(http.StatusOK, "project.html", gin.H{})
	})
	impl.Router.GET("/investors", func(c *gin.Context) {
		c.HTML(http.StatusOK, "investors.html", gin.H{})
	})
	impl.Router.GET("/lk", func(c *gin.Context) {
		c.HTML(http.StatusOK, "lk.html", gin.H{})
	})
	impl.Router.GET("/investments", func(c *gin.Context) {
		c.HTML(http.StatusOK, "investments.html", gin.H{})
	})
	impl.Router.GET("/stages", func(c *gin.Context) {
		c.HTML(http.StatusOK, "stages.html", gin.H{})
	})
	impl.Router.GET("/crm", func(c *gin.Context) {
		c.HTML(http.StatusOK, "crm.html", gin.H{})
	})
	impl.Router.GET("/contracts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contracts.html", gin.H{})
	})
	impl.Router.GET("/expertise", func(c *gin.Context) {
		c.HTML(http.StatusOK, "expertise.html", gin.H{})
	})
	impl.Router.GET("/innovators", func(c *gin.Context) {
		c.HTML(http.StatusOK, "innovators.html", gin.H{})
	})
	impl.Router.GET("/contractors", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contractors.html", gin.H{})
	})
	impl.Router.GET("/risks", func(c *gin.Context) {
		c.HTML(http.StatusOK, "risks.html", gin.H{})
	})
}

func (impl *Implementation) Run() error {
	return impl.Router.Run("localhost:8080")
}
