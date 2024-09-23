package api

import (
	db "github.com/amineadminterraform/go-app/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// serving http
type Server struct {
	router *gin.Engine
	store  db.Store
}

func NewServer(store db.Store) *Server {
	
	server := &Server{store: store}
	router := gin.Default()
	//Creating binding validators
	if v,ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("templateType", validTemplateType)
		v.RegisterValidation("githuburl", validGithubURL)
	}

    // adding route to swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))	
	//adding routes to router
	//Routes for project service
	router.POST("/project", server.CreateProject)
	router.GET("/project/:id", server.GetProject)
	router.GET("/projects", server.ListProjects)
	router.DELETE("/project/:id", server.DeleteProject)
	router.PATCH("/project/:id", server.UpdateProject)
	//Routes for project environment service
	router.POST("/projectenv", server.CreateProjectEnvironment)
	router.GET("/projectenv/:id", server.GetProjectEnvironment)
	router.GET("/projectsenv", server.ListProjectEnvironments)
	router.DELETE("/projectenv/:id", server.DeleteProjectEnvironment)
	router.PATCH("/projectenv/:id", server.UpdateProjectEnvironment)
	//Routes for template service
	router.POST("/template", server.CreateTemplate)
	router.GET("/template/:id", server.GetTemplate)
	router.GET("/templates", server.ListTemplates)
	router.DELETE("/template/:id", server.DeleteTemplate)
	router.PATCH("/template/:id", server.UpdateTemplate)
	//Routes for Argo_workflow service
	router.POST("/argoworkflow", server.CreateArgoWorkflow)
	router.GET("/argoworkflow/:id", server.GetArgoWorkflow)
	router.GET("/argoworkflows", server.ListArgoWorkflows)
	router.DELETE("/argoworkflow/:id", server.DeleteArgoWorkflow)
	router.PATCH("/argoworkflow/:id", server.UpdateArgoWorkflow)
	//Routes for Env_Layers service
	router.POST("/envlayer", server.CreateArgoWorkflow)
	router.GET("/envlayer/:id", server.GetArgoWorkflow)
	router.GET("/envlayers", server.ListArgoWorkflows)
	router.DELETE("/envlayer/:id", server.DeleteArgoWorkflow)
	router.PATCH("/envlayer/:id", server.UpdateArgoWorkflow)
    //Routes for Request service
	router.POST("/request", server.CreateRequest)
	router.GET("/request/:id", server.GetRequest)
	router.GET("/requests", server.ListRequests)
	router.DELETE("/request/:id", server.DeleteRequest)
	router.PATCH("/request/:id", server.UpdateRequest)
	//Routes for Process service
	router.POST("/process", server.CreateProcess)
	router.GET("/process/:id", server.GetProcess)
	router.GET("/processes", server.ListProcesss)
	router.DELETE("/process/:id", server.DeleteProcess)
	router.PATCH("/process/:id", server.UpdateProcess)
	
	
	
	
	
	
	server.router = router
	return server
}

// The error format message from gin
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// The function that starts the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
