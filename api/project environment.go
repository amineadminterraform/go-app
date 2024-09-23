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
type CreateProjectEnvironmentRequest struct {
	GitBranch     string `json:"git_path"`
	Description string `json:"description"`
	ProjectID int64 `json:"project_id"`
}



// CreateProjectEnvironment godoc
// @Summary      	Create a new ProjectEnvironment
// @Description 	Create a ProjectEnvironment and saves it in postgresql
// @Param tags body CreateProjectEnvironmentRequest true "Create ProjectEnvironment"
// @Produce 		application/json
// @host            localhost:8000 
// @ProjectEnvironmentaminos			ProjectEnvironmentaminos
// @Router 	 /projectenv [post]
// @Success 200  {object} SwaggerProjectEnvironment 
// This function will create a new ProjectEnvironment according to the requests
func (server *Server) CreateProjectEnvironment(ctx *gin.Context) {
	var req CreateProjectEnvironmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	arg := db.CreateProjectEnvironmentParams{
		ProjectID: req.ProjectID,
		GitBranch: req.GitBranch,
		Description: pgtype.Text{
			String: req.Description,
			Valid:  true, // Mark as valid if it's not an empty string
		},
	}
	ProjectEnvironment , err := server.store.CreateProjectEnvironment(ctx, arg)
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
	ctx.JSON(http.StatusOK,ProjectEnvironment)
	}


	
	//This structure will create the binding between the uril and the ProjectEnvironment id to get	
	type getProjectEnvironmentRequest struct {
		ID int64 `uri:"id" binding:"required",min:"1"`
	}
// GetProjectEnvironment godoc
// @Summary      	Get a ProjectEnvironment
// @Description 	Get a ProjectEnvironment according to their ID
// @Param        id path int64 true "ProjectEnvironment ID" 
// @Produce 		application/json
// @host            localhost:8000 
// @Tag 			ProjectEnvironment
// @Sucess 200  {object} SwaggerProjectEnvironment
// @Router 		/projectenv/{id} [get]
//This function will get the ProjectEnvironment according to their id	
func (server *Server) GetProjectEnvironment(ctx *gin.Context) {
	var req getProjectEnvironmentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	ProjectEnvironment, err := server.store.GetProjectEnvironment(ctx, req.ID)
	if (err != nil) {
		if pqerr , ok := err.(*pq.Error);ok {
			switch pqerr.Code.Name() {
			case "foreign_key_violation", "unique violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			case "CreateProjectEnvironmentRequest.GitPath":
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
	
	ctx.JSON(http.StatusOK,ProjectEnvironment)
}
//the Structure that will hold the parameters of the list ProjectEnvironment request limit pageSize and offset pageID
type ListProjectEnvironmentRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
// GetProjectEnvironment godoc
// @Summary      	List all ProjectEnvironments
// @Description 	List ProjectEnvironments according to the page_id and page_size
// @Param         page_id query int true "PageID"  // Describing page_id as a query parameter
// @Param         page_size query int true "Page Size"  // Describing page_size as a query parameter
// @Produce 		application/json
// @Tag 			ProjectEnvironment
// @host            localhost:8000 
// @Sucess 200  {[]object} []SwaggerProjectEnvironment
// @Router 		/projectenvs [get]
//This function will get the ProjectEnvironment according to their id	
//The http request to List all the ProjectEnvironments
func (server *Server) ListProjectEnvironments(ctx *gin.Context) {
	var req ListProjectEnvironmentRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.ListProjectEnvironmentsParams{
		Limit: req.PageSize,
		Offset: (req.PageID -1) * req.PageSize,
	}
	ProjectEnvironments, err := server.store.ListProjectEnvironments(ctx, arg) 
		if (err != nil) {
			ctx.JSON(http.StatusInternalServerError,errorResponse(err))
			return 
	}

	ctx.JSON(http.StatusOK,ProjectEnvironments)
	
}

//This structure will create the binding between the uri and the ProjectEnvironment id to delete
type DeleteProjectEnvironmentRequest struct {
	ID int64 `uri:"id" binding: "required",min:"1"`
	}	
//The http request to delete a ProjectEnvironment according to their id in the uril
// GetProjectEnvironment godoc
// @Summary      	delete a ProjectEnvironment with a specified id
// @Description 	Get a ProjectEnvironment from postgresql
// @Param        id path int64 true "ProjectEnvironment ID"
// @Produce 		application/json
// @host            localhost:8000 
// @Tag 			ProjectEnvironment
// @Router 		/projectenv/{id} [delete]
//This function will delete the ProjectEnvironment according to their id	
func (server *Server) DeleteProjectEnvironment(ctx *gin.Context) {
var req getProjectEnvironmentRequest
if err := ctx.ShouldBindUri(&req); err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return 
}
err := server.store.DeleteProjectEnvironment(ctx, req.ID)
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
type UpdateProjectEnvironmentRequest struct {
	ProjectID        int64       `json:"project_id" binding:"required"`
	GitBranch     string      `json:"git_brancg"`
	Description string 		`json:"description"`
}

//This structure will create the binding from the uri to the id
type UpdateProjectEnvironmentUriRequest struct {
    ID int64 `uri:"id" binding:"required"`
}
//The http request to update a ProjectEnvironment according to their id in the uril
// GetProjectEnvironment godoc
// @Summary      	update a ProjectEnvironment with a specified id
// @Description 	update a ProjectEnvironment from postgresql
// @Param tags body UpdateProjectEnvironmentRequest true "update ProjectEnvironment"
// @Param        id path int64 true "ProjectEnvironment ID"
// @Produce 		application/json
// @Tag 			ProjectEnvironment
// @host            localhost:8000 
// @Router 		/projectenv/{id} [patch]
//The http request to update the ProjectEnvironment
func (server *Server) UpdateProjectEnvironment(ctx *gin.Context){
	var req UpdateProjectEnvironmentRequest
	var uriReq UpdateProjectEnvironmentUriRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	if err := ctx.ShouldBindUri(&uriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	arg := db.UpdateProjectEnvironmentParams{
		ID : uriReq.ID,
		GitBranch: req.GitBranch,
		ProjectID: req.ProjectID,
		Description: pgtype.Text{
			String: req.Description,
			Valid:  true, // Mark as valid if it's not an empty string
		},
	}
	
	err := server.store.UpdateProjectEnvironment(ctx, arg)
	if (err != nil) {
		if pqerr , ok := err.(*pq.Error);ok {
			switch pqerr.Code.Name() {
			case "foreign_key_violation", "unique violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			case "CreateProjectEnvironmentRequest.GitPath":
				ctx.JSON(http.StatusBadRequest,errorResponse(err))
			}
			log.Println(pqerr.Code.Name())
		}else {
			ctx.JSON(http.StatusInternalServerError,errorResponse(err))
			return 
		}
	}
	ctx.JSON(http.StatusOK,"")
	} 

