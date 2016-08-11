package main

import (
	"./handler"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()

	handler.WordRoutes(r)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	r.Run(":" + port)
}
