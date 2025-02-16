package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thenopholo/cidade_estado_api.git/internal/app/handlers/locations"
	repositories "github.com/thenopholo/cidade_estado_api.git/internal/infrastructure/repositories/location"
)

func StartServer() {
	e := gin.Default()
	configureRoutes(e)
	err := e.Run(":8080")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server Started")
}

func configureRoutes(e *gin.Engine) {
	locationRepository := repositories.NewLocationRepository()
	locationHandler := locations.NewLocationhanedler(locationRepository)

	g := e.Group("/api/v1")
	{
		g.GET("/states", locationHandler.GetAllStates)
	}
}
