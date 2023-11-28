package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kartheekakkur/service/db/sqlc"
)

type createServiceRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Versions    string `json:"versions" binding:"required"`
}

func (server *Server) createService(ctx *gin.Context) {
	var req createServiceRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	service, err := server.store.GetService(ctx, req.Name)

	//checks if a service with the same name exists in the database
	if service.Name == req.Name {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	//arguments are passed as JSON
	arg := db.CreateServiceParams{
		Name:        req.Name,
		Description: req.Description,
		Versions:    req.Versions,
	}

	service, err = server.store.CreateService(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, service)
}

type getServiceRequest struct {
	Name string `uri:"name" binding:"required"`
}

func (server *Server) getService(ctx *gin.Context) {

	var req getServiceRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	service, err := server.store.GetService(ctx, req.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, service)
}

type listServiceRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=5"`
}

func (server *Server) listService(ctx *gin.Context) {

	var req listServiceRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListServicesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	service, err := server.store.ListServices(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, service)

}

type deleteServiceRequest struct {
	Name string `uri:"name" binding:"required"`
}

func (server *Server) deleteService(ctx *gin.Context) {

	var req deleteServiceRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	service, err := server.store.GetService(ctx, req.Name)

	//checks if a service with the same name exists in the database
	if service.Name != req.Name {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.DeleteService(ctx, req.Name)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, service)
}

type updateServiceRequest struct {
	OldName     string `json:"oldName" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) updateService(ctx *gin.Context) {
	var req updateServiceRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	service, err := server.store.GetService(ctx, req.Name)

	//checks if a service with the same name exists in the database
	if service.Name == req.Name {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	//arguments are passed as JSON
	arg := db.UpdateServiceParams{
		Name:        req.OldName,
		Description: req.Description,
		Name_2:      req.Name,
	}
	service, err = server.store.UpdateService(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, service)
}
