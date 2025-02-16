package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	g := e.Group("/api/v1")
	{
		g.GET("/states", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"hello": "World",
			})
		})
	}
}
