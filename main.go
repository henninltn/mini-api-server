package main

import (
	"./handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	handler.ThreadRoutes(router)
	handler.ResponseRoutes(router)

	router.Run(":8080")
}
