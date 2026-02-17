package context

import (
	"openingjobs/internal/opening"

	"github.com/gin-gonic/gin"
)

func InitializeHandlerContexts() error {
	// Initilize router
	router := gin.Default()

	//Initialize contexts V1
	v1 := router.Group("/api/v1")
	opening.InitializeOpeningHandlerConext(v1)

	//Run the server
	router.Run()

	return nil
}
