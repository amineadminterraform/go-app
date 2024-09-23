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
type CreateRequestRequest struct {
	Source           string             `json:"source"`
	Payload           string             `json:"payload"`
	LayerID        int64              `json:"layer_id"`
	Status           string             `json:"status"`
}



// CreateRequest godoc
// @Summary      	Create a new Request
// @Description 	Create a Request and saves it in postgresql
// @Param tags body CreateRequestRequest true "Create Request"
// @Produce 		application/json
// @host            localhost:8000 
// @Requestaminos			Requestaminos
// @Router 	 /request [post]
// @Success 200  {object} SwaggerRequest 
// This function will create a new Request according to the requests
func (server *Server) CreateRequest(ctx *gin.Context) {
	var req CreateRequestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	arg := db.CreateRequestParams{
		LayerID: req.LayerID,
		Source: req.Source,
		Payload: req.Payload,
		Status: req.Status,
		}
	Request , err := server.store.CreateRequest(ctx, arg)
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
	ctx.JSON(http.StatusOK,Request)
	}


	
	//This structure will create the binding between the uril and the Request id to get	
	type getRequestRequest struct {
		ID int64 `uri:"id" binding:"required",min:"1"`
	}
// GetRequest godoc
// @Summary      	Get a Request
// @Description 	Get a Request according to their ID
// @Param        id path int64 true "Request ID" 
// @Produce 		application/json
// @host            localhost:8000 
// @Tag 			Request
// @Sucess 200  {object} SwaggerRequest
// @Router 		/request/{id} [get]
//This function will get the Request according to their id	

func (server *Server) GetRequest(ctx *gin.Context) {
	var req getRequestRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	Request, err := server.store.GetRequest(ctx, req.ID)
	if (err != nil) {
		if pqerr , ok := err.(*pq.Error);ok {
			switch pqerr.Code.Name() {
			case "foreign_key_violation", "unique violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			case "CreateRequestRequest.GitPath":
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
	
	ctx.JSON(http.StatusOK,Request)
}
//the Structure that will hold the parameters of the list Request request limit pageSize and offset pageID
type ListRequestRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
// GetRequest godoc
// @Summary      	List all Requests
// @Description 	List Requests according to the page_id and page_size
// @Param         page_id query int true "PageID"  // Describing page_id as a query parameter
// @Param         page_size query int true "Page Size"  // Describing page_size as a query parameter
// @Produce 		application/json
// @Tag 			Request
// @host            localhost:8000 
// @Sucess 200  {[]object} []SwaggerRequest
// @Router 		/requests [get]
//This function will get the Request according to their id	
//The http request to List all the Requests
func (server *Server) ListRequests(ctx *gin.Context) {
	var req ListRequestRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.ListRequestsParams{
		Limit: req.PageSize,
		Offset: (req.PageID -1) * req.PageSize,
	}
	Requests, err := server.store.ListRequests(ctx, arg) 
		if (err != nil) {
			ctx.JSON(http.StatusInternalServerError,errorResponse(err))
			return 
	}

	ctx.JSON(http.StatusOK,Requests)
	
}

//This structure will create the binding between the uri and the Request id to delete
type DeleteRequestRequest struct {
	ID int64 `uri:"id" binding: "required",min:"1"`
	}	
//The http request to delete a Request according to their id in the uril
// GetRequest godoc
// @Summary      	delete a Request with a specified id
// @Description 	Get a Request from postgresql
// @Param        id path int64 true "Request ID"
// @Produce 		application/json
// @host            localhost:8000 
// @Tag 			Request
// @Router 		/request/{id} [delete]
//This function will delete the Request according to their id	
func (server *Server) DeleteRequest(ctx *gin.Context) {
var req getRequestRequest
if err := ctx.ShouldBindUri(&req); err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return 
}
err := server.store.DeleteRequest(ctx, req.ID)
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
type UpdateRequestRequest struct {
	Source           string             `json:"source"`
	Payload           string             `json:"payload"`
	LayerID        int64              `json:"layer_id"`
	Status           string             `json:"status"`

}

//This structure will create the binding from the uri to the id
type UpdateRequestUriRequest struct {
    ID int64 `uri:"id" binding:"required"`
}
//The http request to update a Request according to their id in the uril
// GetRequest godoc
// @Summary      	update a Request with a specified id
// @Description 	update a Request from postgresql
// @Param tags body UpdateRequestRequest true "update Request"
// @Param        id path int64 true "Request ID"
// @Produce 		application/json
// @Tag 			Request
// @host            localhost:8000 
// @Router 		/request/{id} [patch]
//The http request to update the Request
func (server *Server) UpdateRequest(ctx *gin.Context){
	var req UpdateRequestRequest
	var uriReq UpdateRequestUriRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	if err := ctx.ShouldBindUri(&uriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	arg := db.UpdateRequestParams{
		ID : uriReq.ID,
		LayerID: req.LayerID,
		Source: req.Source,
		Payload: req.Payload,
		Status: req.Status,
	}
	
	err := server.store.UpdateRequest(ctx, arg)
	if (err != nil) {
		if pqerr , ok := err.(*pq.Error);ok {
			switch pqerr.Code.Name() {
			case "foreign_key_violation", "unique violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			case "CreateRequestRequest.GitPath":
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

