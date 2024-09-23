package api

import (
	"database/sql"
	"log"
	"net/http"

	db "github.com/amineadminterraform/go-app/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lib/pq"
)

// This structure will create the necessary bindings between the body of the request and the name , gitpath and description
type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	GitPath     string `json:"git_path" binding:"githuburl"`
	Description string `json:"description"`
}




// CreateProject godoc
// @Summary      	Create a new project
// @Description 	Create a project and saves it in postgresql
// @Param tags body CreateProjectRequest true "Create project"
// @Produce 		application/json
// @host            localhost:8000 
// @Projectaminos			projectaminos
// @Router 	 /project [post]
// @Success 200  {object} SwaggerProject 
// This function will create a new project according to the requests
func (server *Server) CreateProject(ctx *gin.Context) {
	var req CreateProjectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	arg := db.CreateProjectParams{
		Name: req.Name,
		GitPath: req.GitPath,
		Description: pgtype.Text{
			String: req.Description,
			Valid:  true, // Mark as valid if it's not an empty string
		},
	}
	project , err := server.store.CreateProject(ctx, arg)
	if (err != nil) {
		if pqerr , ok := err.(*pq.Error);ok {
			switch pqerr.Code.Name() {
			case "foreign_key_violation", "unique violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			}
			log.Println(pqerr.Code.Name())
		}else {
			ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		}
		return 
			
		
	}
	ctx.JSON(http.StatusOK,project)
	}


	
	//This structure will create the binding between the uril and the project id to get	
	type getProjectRequest struct {
		ID int64 `uri:"id" binding:"required",min:"1"`
	}
// GetProject godoc
// @Summary      	Get a project
// @Description 	Get a project according to their ID
// @Param        id path int64 true "Project ID" 
// @Produce 		application/json
// @host            localhost:8000 
// @Tag 			project
// @Sucess 200  {object} SwaggerProject
// @Router 		/project/{id} [get]
//This function will get the project according to their id	
func (server *Server) GetProject(ctx *gin.Context) {
	var req getProjectRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	project, err := server.store.GetProject(ctx, req.ID)
	if (err != nil) {
		if pqerr , ok := err.(*pq.Error);ok {
			switch pqerr.Code.Name() {
			case "foreign_key_violation", "unique violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			case "CreateProjectRequest.GitPath":
				ctx.JSON(http.StatusBadRequest,errorResponse(err))
			}

			log.Println(pqerr.Code.Name())
		}

		if (err == sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound,errorResponse(err))
			return 
			
		}
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK,project)
}
//the Structure that will hold the parameters of the list project request limit pageSize and offset pageID
type ListProjectRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
// GetProject godoc
// @Summary      	List all projects
// @Description 	List projects according to the page_id and page_size
// @Param         page_id query int true "PageID"  // Describing page_id as a query parameter
// @Param         page_size query int true "Page Size"  // Describing page_size as a query parameter
// @Produce 		application/json
// @Tag 			project
// @host            localhost:8000 
// @Sucess 200  {[]object} []SwaggerProject
// @Router 		/projects [get]
//This function will get the project according to their id	
//The http request to List all the projects
func (server *Server) ListProjects(ctx *gin.Context) {
	var req ListProjectRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.ListProjectsParams{
		Limit: req.PageSize,
		Offset: (req.PageID -1) * req.PageSize,
	}
	projects, err := server.store.ListProjects(ctx, arg) 
		if (err != nil) {
			ctx.JSON(http.StatusInternalServerError,errorResponse(err))
			return 
	}

	ctx.JSON(http.StatusOK,projects)
	
}

//This structure will create the binding between the uri and the project id to delete
type DeleteProjectRequest struct {
	ID int64 `uri:"id" binding:"required",min:"1"`
	}	
//The http request to delete a project according to their id in the uril
// GetProject godoc
// @Summary      	delete a project with a specified id
// @Description 	Get a project from postgresql
// @Param        id path int64 true "Project ID"
// @Produce 		application/json
// @host            localhost:8000 
// @Tag 			project
// @Router 		/project/{id} [delete]
//This function will delete the project according to their id	
func (server *Server) DeleteProject(ctx *gin.Context) {
var req getProjectRequest
if err := ctx.ShouldBindUri(&req); err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return 
}
err := server.store.DeleteProject(ctx, req.ID)
if (err != nil) {
	if (err == sql.ErrNoRows) {
		ctx.JSON(http.StatusNotFound,errorResponse(err))
		return 
		
	}
	ctx.JSON(http.StatusInternalServerError,errorResponse(err))
	return
}

ctx.JSON(http.StatusAccepted,nil)
}


//This structure will create the json binding from the body of the request to update name , gitpath and description
type UpdateProjectRequest struct {
	Name        string      `json:"name" binding:"required"`
	GitPath     string      `json:"git_path" binding:"githuburl"`
	Description string 		`json:"description"`
}

//This structure will create the binding from the uri to the id
type UpdateProjectUriRequest struct {
    ID int64 `uri:"id" binding:"required"`
}
//The http request to update a project according to their id in the uril
// GetProject godoc
// @Summary      	update a project with a specified id
// @Description 	update a project from postgresql
// @Param tags body UpdateProjectRequest true "update project"
// @Param        id path int64 true "Project ID"
// @Produce 		application/json
// @Tag 			project
// @host            localhost:8000 
// @Router 		/project/{id} [patch]
//The http request to update the project
func (server *Server) UpdateProject(ctx *gin.Context){
	var req UpdateProjectRequest
	var uriReq UpdateProjectUriRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	if err := ctx.ShouldBindUri(&uriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	arg := db.UpdateProjectParams{
		ID : uriReq.ID,
		Name: req.Name,
		GitPath: req.GitPath,
		Description: pgtype.Text{
			String: req.Description,
			Valid:  true, // Mark as valid if it's not an empty string
		},
	}
	project , err := server.store.UpdateProject(ctx, arg)
	if (err != nil) {
		if pqerr , ok := err.(*pq.Error);ok {
			switch pqerr.Code.Name() {
			case "foreign_key_violation", "unique violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			case "CreateProjectRequest.GitPath":
				ctx.JSON(http.StatusBadRequest,errorResponse(err))
			}
			log.Println(pqerr.Code.Name())
		}else {
			ctx.JSON(http.StatusInternalServerError,errorResponse(err))
			return 
		}
	}
	ctx.JSON(http.StatusOK,project)
	} 

