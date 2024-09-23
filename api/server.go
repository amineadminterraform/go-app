package api

import (
	db "github.com/amineadminterraform/go-app/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Queries struct {
	db *pgx.Conn
}

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
	router.POST("/project", server.CreateProject)
	router.GET("/project/:id", server.GetProject)
	router.GET("/projects", server.ListProjects)
	router.DELETE("/project/:id", server.DeleteProject)
	router.PATCH("/project/:id", server.UpdateProject)

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
