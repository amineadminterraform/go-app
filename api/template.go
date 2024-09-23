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
type CreateTemplateRequest struct {
	Name        string `json:"name" binding:"required"`
	Path     string `json:"git_path"`
	Type string `json:"type"`
}



// CreateTemplate godoc
// @Summary      	Create a new Template
// @Description 	Create a Template and saves it in postgresql
// @Param tags body CreateTemplateRequest true "Create Template"
// @Produce 		application/json
// @host            localhost:8000 
// @Templateaminos			Templateaminos
// @Router 	 /template [post]
// @Success 200  {object} SwaggerTemplate 
// This function will create a new Template according to the requests
func (server *Server) CreateTemplate(ctx *gin.Context) {
	var req CreateTemplateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	arg := db.CreateTemplateParams{
		Name: req.Name,
		Path: req.Path,
		Type: req.Type,
	}
	Template , err := server.store.CreateTemplate(ctx, arg)
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
	ctx.JSON(http.StatusOK,Template)
	}


	
	//This structure will create the binding between the uril and the Template id to get	
	type getTemplateRequest struct {
		ID int64 `uri:"id" binding:"required",min:"1"`
	}
// GetTemplate godoc
// @Summary      	Get a Template
// @Description 	Get a Template according to their ID
// @Param        id path int64 true "Template ID" 
// @Produce 		application/json
// @host            localhost:8000 
// @Tag 			Template
// @Sucess 200  {object} SwaggerTemplate
// @Router 		/template/{id} [get]
//This function will get the Template according to their id	
func (server *Server) GetTemplate(ctx *gin.Context) {
	var req getTemplateRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	Template, err := server.store.GetTemplate(ctx, req.ID)
	if (err != nil) {
		if pqerr , ok := err.(*pq.Error);ok {
			switch pqerr.Code.Name() {
			case "foreign_key_violation", "unique violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			case "CreateTemplateRequest.GitPath":
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
	
	ctx.JSON(http.StatusOK,Template)
}
//the Structure that will hold the parameters of the list Template request limit pageSize and offset pageID
type ListTemplateRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
// GetTemplate godoc
// @Summary      	List all Templates
// @Description 	List Templates according to the page_id and page_size
// @Param         page_id query int true "PageID"  // Describing page_id as a query parameter
// @Param         page_size query int true "Page Size"  // Describing page_size as a query parameter
// @Produce 		application/json
// @Tag 			Template
// @host            localhost:8000 
// @Sucess 200  {[]object} []SwaggerTemplate
// @Router 		/templates [get]
//This function will get the Template according to their id	
//The http request to List all the Templates
func (server *Server) ListTemplates(ctx *gin.Context) {
	var req ListTemplateRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.ListTemplatesParams{
		Limit: req.PageSize,
		Offset: (req.PageID -1) * req.PageSize,
	}
	Templates, err := server.store.ListTemplates(ctx, arg) 
		if (err != nil) {
			ctx.JSON(http.StatusInternalServerError,errorResponse(err))
			return 
	}

	ctx.JSON(http.StatusOK,Templates)
	
}

//This structure will create the binding between the uri and the Template id to delete
type DeleteTemplateRequest struct {
	ID int64 `uri:"id" binding:"required",min:"1"`
	}	
//The http request to delete a Template according to their id in the uril
// GetTemplate godoc
// @Summary      	delete a Template with a specified id
// @Description 	Get a Template from postgresql
// @Param        id path int64 true "Template ID"
// @Produce 		application/json
// @host            localhost:8000 
// @Tag 			Template
// @Router 		/template/{id} [delete]
//This function will delete the Template according to their id	
func (server *Server) DeleteTemplate(ctx *gin.Context) {
var req getTemplateRequest
if err := ctx.ShouldBindUri(&req); err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return 
}
err := server.store.DeleteTemplate(ctx, req.ID)
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
type UpdateTemplateRequest struct {
	Name        string      `json:"name" binding:"required"`
	Path     string      `json:"git_path"`
	Type string 		`json:"type"`
}

//This structure will create the binding from the uri to the id
type UpdateTemplateUriRequest struct {
    ID int64 `uri:"id" binding:"required"`
}
//The http request to update a Template according to their id in the uril
// GetTemplate godoc
// @Summary      	update a Template with a specified id
// @Description 	update a Template from postgresql
// @Param tags body UpdateTemplateRequest true "update Template"
// @Param        id path int64 true "Template ID"
// @Produce 		application/json
// @Tag 			Template
// @host            localhost:8000 
// @Router 		/template/{id} [patch]
//The http request to update the Template
func (server *Server) UpdateTemplate(ctx *gin.Context){
	var req UpdateTemplateRequest
	var uriReq UpdateTemplateUriRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	if err := ctx.ShouldBindUri(&uriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
	arg := db.UpdateTemplateParams{
		ID : uriReq.ID,
		Name: req.Name,
		Path: req.Path,
		Type: req.Type,
	}
	err := server.store.UpdateTemplate(ctx, arg)
	if (err != nil) {
		if pqerr , ok := err.(*pq.Error);ok {
			switch pqerr.Code.Name() {
			case "foreign_key_violation", "unique violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			case "CreateTemplateRequest.GitPath":
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

