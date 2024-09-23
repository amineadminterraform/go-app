package api

import (
	"database/sql"
	"log"
	"net/http"

	db "github.com/amineadminterraform/go-app/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// This structure will create the necessary bindings between the body of the request and the name , gitpath and description
type CreateProcessRequest struct {
	ArgoID     int64              `json:"argo_id"`
	Name       string             `json:"name"`
	TemplateID int64              `json:"template_id"`
}



// CreateProcess godoc
// @Summary      	Create a new Process
// @Description 	Create a Process and saves it in postgresql
// @Param tags body CreateProcessRequest true "Create Process"
// @Produce 		application/json
// @host            localhost:8000 
// @Processaminos			Processaminos
// @Router 	 /process [post]
// @Success 200  {object} SwaggerProcess 
// This function will create a new Process according to the requests
func (server *Server) CreateProcess(ctx *gin.Context) {
	var req CreateProcessRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	arg := db.CreateProcessParams{
		ArgoID: req.ArgoID,
		Name: req.Name,
		TemplateID: req.TemplateID,
		}
	Process , err := server.store.CreateProcess(ctx, arg)
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
	ctx.JSON(http.StatusOK,Process)
	}


	
	//This structure will create the binding between the uril and the Process id to get	
	type getProcessRequest struct {
		ID int64 `uri:"id" binding:"required",min:"1"`
	}
// GetProcess godoc
// @Summary      	Get a Process
// @Description 	Get a Process according to their ID
// @Param        id path int64 true "Process ID" 
// @Produce 		application/json
// @host            localhost:8000 
// @Tag 			Process
// @Sucess 200  {object} SwaggerProcess
// @Router 		/process/{id} [get]
//This function will get the Process according to their id	

func (server *Server) GetProcess(ctx *gin.Context) {
	var req getProcessRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	Process, err := server.store.GetProcess(ctx, req.ID)
	if (err != nil) {
		if pqerr , ok := err.(*pq.Error);ok {
			switch pqerr.Code.Name() {
			case "foreign_key_violation", "unique violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			case "CreateProcessRequest.GitPath":
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
	
	ctx.JSON(http.StatusOK,Process)
}
//the Structure that will hold the parameters of the list Process request limit pageSize and offset pageID
type ListProcessRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
// GetProcess godoc
// @Summary      	List all Processs
// @Description 	List Processs according to the page_id and page_size
// @Param         page_id query int true "PageID"  // Describing page_id as a query parameter
// @Param         page_size query int true "Page Size"  // Describing page_size as a query parameter
// @Produce 		application/json
// @Tag 			Process
// @host            localhost:8000 
// @Sucess 200  {[]object} []SwaggerProcess
// @Router 		/processes [get]
//This function will get the Process according to their id	
//The http request to List all the Processs
func (server *Server) ListProcesss(ctx *gin.Context) {
	var req ListProcessRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.ListProcesssParams{
		Limit: req.PageSize,
		Offset: (req.PageID -1) * req.PageSize,
	}
	Processs, err := server.store.ListProcesss(ctx, arg) 
		if (err != nil) {
			ctx.JSON(http.StatusInternalServerError,errorResponse(err))
			return 
	}

	ctx.JSON(http.StatusOK,Processs)
	
}

//This structure will create the binding between the uri and the Process id to delete
type DeleteProcessRequest struct {
	ID int64 `uri:"id" binding: "required",min:"1"`
	}	
//The http request to delete a Process according to their id in the uril
// GetProcess godoc
// @Summary      	delete a Process with a specified id
// @Description 	Get a Process from postgresql
// @Param        id path int64 true "Process ID"
// @Produce 		application/json
// @host            localhost:8000 
// @Tag 			Process
// @Router 		/process/{id} [delete]
//This function will delete the Process according to their id	
func (server *Server) DeleteProcess(ctx *gin.Context) {
var req getProcessRequest
if err := ctx.ShouldBindUri(&req); err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return 
}
err := server.store.DeleteProcess(ctx, req.ID)
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
type UpdateProcessRequest struct {
	ArgoID     int64              `json:"argo_id"`
	Name       string             `json:"name"`
	TemplateID int64              `json:"template_id"`

}

//This structure will create the binding from the uri to the id
type UpdateProcessUriRequest struct {
    ID int64 `uri:"id" binding:"required"`
}
//The http request to update a Process according to their id in the uril
// GetProcess godoc
// @Summary      	update a Process with a specified id
// @Description 	update a Process from postgresql
// @Param tags body UpdateProcessRequest true "update Process"
// @Param        id path int64 true "Process ID"
// @Produce 		application/json
// @Tag 			Process
// @host            localhost:8000 
// @Router 		/process/{id} [patch]
//The http request to update the Process
func (server *Server) UpdateProcess(ctx *gin.Context){
	var req UpdateProcessRequest
	var uriReq UpdateProcessUriRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	if err := ctx.ShouldBindUri(&uriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	arg := db.UpdateProcessParams{
		ID : uriReq.ID,
		ArgoID: req.ArgoID,
		Name: req.Name,
		TemplateID: req.TemplateID,
	}
	
	err := server.store.UpdateProcess(ctx, arg)
	if (err != nil) {
		if pqerr , ok := err.(*pq.Error);ok {
			switch pqerr.Code.Name() {
			case "foreign_key_violation", "unique violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			case "CreateProcessRequest.GitPath":
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

