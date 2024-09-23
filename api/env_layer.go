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
type CreateEnvLayerRequest struct {
	S3Path           string             `json:"s3_path"`
	ProcessID        int64              `json:"process_id"`
	CurrentRequestID pgtype.Int8        `json:"current_request_id"`
	EnvironmentID int64 `json:"environment_id"`
}




// CreateEnvLayer godoc
// @Summary      	Create a new EnvLayer
// @Description 	Create a EnvLayer and saves it in postgresql
// @Param tags body CreateEnvLayerRequest true "Create EnvLayer"
// @Produce 		application/json
// @host            localhost:8000 
// @EnvLayeraminos			EnvLayeraminos
// @Router 	 /envlayer [post]
// @Success 200  {object} SwaggerEnvLayer 
// This function will create a new EnvLayer according to the requests
func (server *Server) CreateEnvLayer(ctx *gin.Context) {
	var req CreateEnvLayerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	arg := db.CreateEnvLayerParams{
		EnvironmentID: req.EnvironmentID,
		S3Path: req.S3Path,
		ProcessID: req.ProcessID,
		CurrentRequestID: req.CurrentRequestID,
		}
	EnvLayer , err := server.store.CreateEnvLayer(ctx, arg)
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
	ctx.JSON(http.StatusOK,EnvLayer)
	}


	
	//This structure will create the binding between the uril and the EnvLayer id to get	
	type getEnvLayerRequest struct {
		ID int64 `uri:"id" binding:"required",min:"1"`
	}
// GetEnvLayer godoc
// @Summary      	Get a EnvLayer
// @Description 	Get a EnvLayer according to their ID
// @Param        id path int64 true "EnvLayer ID" 
// @Produce 		application/json
// @host            localhost:8000 
// @Tag 			EnvLayer
// @Sucess 200  {object} SwaggerEnvLayer
// @Router 		/envlayer/{id} [get]
//This function will get the EnvLayer according to their id	

func (server *Server) GetEnvLayer(ctx *gin.Context) {
	var req getEnvLayerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	EnvLayer, err := server.store.GetEnvLayer(ctx, req.ID)
	if (err != nil) {
		if pqerr , ok := err.(*pq.Error);ok {
			switch pqerr.Code.Name() {
			case "foreign_key_violation", "unique violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			case "CreateEnvLayerRequest.GitPath":
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
	
	ctx.JSON(http.StatusOK,EnvLayer)
}
//the Structure that will hold the parameters of the list EnvLayer request limit pageSize and offset pageID
type ListEnvLayerRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
// GetEnvLayer godoc
// @Summary      	List all EnvLayers
// @Description 	List EnvLayers according to the page_id and page_size
// @Param         page_id query int true "PageID"  // Describing page_id as a query parameter
// @Param         page_size query int true "Page Size"  // Describing page_size as a query parameter
// @Produce 		application/json
// @Tag 			EnvLayer
// @host            localhost:8000 
// @Sucess 200  {[]object} []SwaggerEnvLayer
// @Router 		/envlayers [get]
//This function will get the EnvLayer according to their id	
//The http request to List all the EnvLayers
func (server *Server) ListEnvLayers(ctx *gin.Context) {
	var req ListEnvLayerRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.ListEnvLayersParams{
		Limit: req.PageSize,
		Offset: (req.PageID -1) * req.PageSize,
	}
	EnvLayers, err := server.store.ListEnvLayers(ctx, arg) 
		if (err != nil) {
			ctx.JSON(http.StatusInternalServerError,errorResponse(err))
			return 
	}

	ctx.JSON(http.StatusOK,EnvLayers)
	
}

//This structure will create the binding between the uri and the EnvLayer id to delete
type DeleteEnvLayerRequest struct {
	ID int64 `uri:"id" binding: "required",min:"1"`
	}	
//The http request to delete a EnvLayer according to their id in the uril
// GetEnvLayer godoc
// @Summary      	delete a EnvLayer with a specified id
// @Description 	Get a EnvLayer from postgresql
// @Param        id path int64 true "EnvLayer ID"
// @Produce 		application/json
// @host            localhost:8000 
// @Tag 			EnvLayer
// @Router 		/envlayer/{id} [delete]
//This function will delete the EnvLayer according to their id	
func (server *Server) DeleteEnvLayer(ctx *gin.Context) {
var req getEnvLayerRequest
if err := ctx.ShouldBindUri(&req); err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return 
}
err := server.store.DeleteEnvLayer(ctx, req.ID)
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
type UpdateEnvLayerRequest struct {
	S3Path           string             `json:"s3_path"`
	ProcessID        int64              `json:"process_id"`
	CurrentRequestID pgtype.Int8        `json:"current_request_id"`
	EnvironmentID int64 `json:"environment_id"`
}

//This structure will create the binding from the uri to the id
type UpdateEnvLayerUriRequest struct {
    ID int64 `uri:"id" binding:"required"`
}
//The http request to update a EnvLayer according to their id in the uril
// GetEnvLayer godoc
// @Summary      	update a EnvLayer with a specified id
// @Description 	update a EnvLayer from postgresql
// @Param tags body UpdateEnvLayerRequest true "update EnvLayer"
// @Param        id path int64 true "EnvLayer ID"
// @Produce 		application/json
// @Tag 			EnvLayer
// @host            localhost:8000 
// @Router 		/envlayer/{id} [patch]
//The http request to update the EnvLayer
func (server *Server) UpdateEnvLayer(ctx *gin.Context){
	var req UpdateEnvLayerRequest
	var uriReq UpdateEnvLayerUriRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	if err := ctx.ShouldBindUri(&uriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	arg := db.UpdateEnvLayerParams{
		ID : uriReq.ID,
		S3Path: req.S3Path,
		ProcessID: req.ProcessID,
		CurrentRequestID: req.CurrentRequestID,
	}
	
	err := server.store.UpdateEnvLayer(ctx, arg)
	if (err != nil) {
		if pqerr , ok := err.(*pq.Error);ok {
			switch pqerr.Code.Name() {
			case "foreign_key_violation", "unique violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			case "CreateEnvLayerRequest.GitPath":
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

