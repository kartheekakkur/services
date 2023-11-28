package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/kartheekakkur/service/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// New Server instance
func NewServer(store *db.Store) *Server {

	server := &Server{store: store}
	router := gin.Default()

	//add routes to perform CRUD.
	router.POST("/service", server.createService)              // creates a new service after validating if the service with a name exists
	router.GET("/service/:name", server.getService)            //searches for a service with specific name and passed as URI
	router.GET("/service", server.listService)                 // returns the list of services sorted by ID taking in page_id and page_size parameters.
	router.POST("/service/delete/:name", server.deleteService) // deletes service
	router.POST("/service/update", server.updateService)       // update service

	server.router = router
	return server

}

func (server *Server) Start(address string) error {

	return server.router.Run(address)

}
func errorResponse(err error) gin.H {
	return gin.H{"error": "error"}
}
