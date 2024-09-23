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
type CreateArgoWorkflowRequest struct {
	Name        string `json:"name" binding:"required"`
	Path     string `json:"path"`
	Description string `json:"description"`
}


// CreateArgoWorkflow godoc
// @Summary      	Create a new ArgoWorkflow
// @Description 	Create a ArgoWorkflow and saves it in postgresql
// @Param tags body CreateArgoWorkflowRequest true "Create ArgoWorkflow"
// @Produce 		application/json
// @host            localhost:8000 
// @ArgoWorkflowaminos			ArgoWorkflowaminos
// @Router 	 /argoworkflow [post]
// @Success 200  {object} SwaggerArgoWorkflow 
// This function will create a new ArgoWorkflow according to the requests
func (server *Server) CreateArgoWorkflow(ctx *gin.Context) {
	var req CreateArgoWorkflowRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	arg := db.CreateArgoWorkflowParams{
		Name: req.Name,
		Path: req.Path,
		Description: pgtype.Text{
			String: req.Description,
			Valid:  true, // Mark as valid if it's not an empty string
		},
	}
	ArgoWorkflow , err := server.store.CreateArgoWorkflow(ctx, arg)
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
	ctx.JSON(http.StatusOK,ArgoWorkflow)
	}


	
	//This structure will create the binding between the uril and the ArgoWorkflow id to get	
	type getArgoWorkflowRequest struct {
		ID int64 `uri:"id" binding:"required",min:"1"`
	}
// GetArgoWorkflow godoc
// @Summary      	Get a ArgoWorkflow
// @Description 	Get a ArgoWorkflow according to their ID
// @Param        id path int64 true "ArgoWorkflow ID" 
// @Produce 		application/json
// @host            localhost:8000 
// @Tag 			ArgoWorkflow
// @Sucess 200  {object} SwaggerArgoWorkflow
// @Router 		/argoworkflow/{id} [get]
//This function will get the ArgoWorkflow according to their id	
func (server *Server) GetArgoWorkflow(ctx *gin.Context) {
	var req getArgoWorkflowRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	ArgoWorkflow, err := server.store.GetArgoWorkflow(ctx, req.ID)
	if (err != nil) {
		if pqerr , ok := err.(*pq.Error);ok {
			switch pqerr.Code.Name() {
			case "foreign_key_violation", "unique violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			case "CreateArgoWorkflowRequest.GitPath":
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
	
	ctx.JSON(http.StatusOK,ArgoWorkflow)
}
//the Structure that will hold the parameters of the list ArgoWorkflow request limit pageSize and offset pageID
type ListArgoWorkflowRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
// GetArgoWorkflow godoc
// @Summary      	List all ArgoWorkflows
// @Description 	List ArgoWorkflows according to the page_id and page_size
// @Param         page_id query int true "PageID"  // Describing page_id as a query parameter
// @Param         page_size query int true "Page Size"  // Describing page_size as a query parameter
// @Produce 		application/json
// @Tag 			ArgoWorkflow
// @host            localhost:8000 
// @Sucess 200  {[]object} []SwaggerArgoWorkflow
// @Router 		/argoworkflows [get]
//This function will get the ArgoWorkflow according to their id	
//The http request to List all the ArgoWorkflows
func (server *Server) ListArgoWorkflows(ctx *gin.Context) {
	var req ListArgoWorkflowRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.ListArgoWorkflowsParams{
		Limit: req.PageSize,
		Offset: (req.PageID -1) * req.PageSize,
	}
	ArgoWorkflows, err := server.store.ListArgoWorkflows(ctx, arg) 
		if (err != nil) {
			ctx.JSON(http.StatusInternalServerError,errorResponse(err))
			return 
	}

	ctx.JSON(http.StatusOK,ArgoWorkflows)
	
}

//This structure will create the binding between the uri and the ArgoWorkflow id to delete
type DeleteArgoWorkflowRequest struct {
	ID int64 `uri:"id" binding:"required",min:"1"`
	}	
//The http request to delete a ArgoWorkflow according to their id in the uril
// GetArgoWorkflow godoc
// @Summary      	delete a ArgoWorkflow with a specified id
// @Description 	Get a ArgoWorkflow from postgresql
// @Param        id path int64 true "ArgoWorkflow ID"
// @Produce 		application/json
// @host            localhost:8000 
// @Tag 			ArgoWorkflow
// @Router 		/argoworkflow/{id} [delete]
//This function will delete the ArgoWorkflow according to their id	
func (server *Server) DeleteArgoWorkflow(ctx *gin.Context) {
var req getArgoWorkflowRequest
if err := ctx.ShouldBindUri(&req); err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return 
}
err := server.store.DeleteArgoWorkflow(ctx, req.ID)
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
type UpdateArgoWorkflowRequest struct {
	Name        string      `json:"name" binding:"required"`
	Path     string      `json:"git_path"`
	Description string 		`json:"description"`

}

//This structure will create the binding from the uri to the id
type UpdateArgoWorkflowUriRequest struct {
    ID int64 `uri:"id" binding:"required"`
}
//The http request to update a ArgoWorkflow according to their id in the uril
// GetArgoWorkflow godoc
// @Summary      	update a ArgoWorkflow with a specified id
// @Description 	update a ArgoWorkflow from postgresql
// @Param tags body UpdateArgoWorkflowRequest true "update ArgoWorkflow"
// @Param        id path int64 true "ArgoWorkflow ID"
// @Produce 		application/json
// @Tag 			ArgoWorkflow
// @host            localhost:8000 
// @Router 		/argoworkflow/{id} [patch]
//The http request to update the ArgoWorkflow
func (server *Server) UpdateArgoWorkflow(ctx *gin.Context){
	var req UpdateArgoWorkflowRequest
	var uriReq UpdateArgoWorkflowUriRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	if err := ctx.ShouldBindUri(&uriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	arg := db.UpdateArgoWorkflowParams{
		ID : uriReq.ID,
		Name: req.Name,
		Path: req.Path,
		Description: pgtype.Text{
			String: req.Description,
			Valid:  true, // Mark as valid if it's not an empty string
		},
	}
	err := server.store.UpdateArgoWorkflow(ctx, arg)
	if (err != nil) {
		if pqerr , ok := err.(*pq.Error);ok {
			switch pqerr.Code.Name() {
			case "foreign_key_violation", "unique violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			case "CreateArgoWorkflowRequest.GitPath":
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

