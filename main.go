package main

import (
	"./handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	handler.ThreadRoutes(router)

	router.Run(":8080")
}
