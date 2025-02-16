package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	e := gin.Default()
	err := e.Run(":8080")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server Started")
}
