package api

import (
	"database/sql"
	"net/http"

	db "github.com/amineadminterraform/go-app/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateProjectRequest struct {
	Name        string      `json:"name" binding:"required"`
	GitPath     string      `json:"git_path"`
	Description string 		`json:"description"`
}
func (server *Server) CreateProject(ctx *gin.Context){
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
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return 
	}
	ctx.JSON(http.StatusOK,project)
	} 
	type getProjectRequest struct {
		ID int64 `uri:"id" binding:"required",min:"1"`
	}	
	func (server *Server) GetProject(ctx *gin.Context) {
	var req getProjectRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	project, err := server.store.GetProject(ctx, req.ID)
	if (err != nil) {
		if (err == sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound,errorResponse(err))
			return 
			
		}
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK,project)
}
type ListProjectRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
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
type DeleteProjectRequest struct {
	ID int64 `uri:"id" binding:"required",min:"1"`
}	
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

type UpdateProjectRequest struct {
	Name        string      `json:"name" binding:"required"`
	GitPath     string      `json:"git_path"`
	Description string 		`json:"description"`
}
type UpdateProjectUriRequest struct {
    ID int64 `uri:"id" binding:"required"`
}
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
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return 
	}
	ctx.JSON(http.StatusOK,project)
	} 
