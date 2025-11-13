package main

import (
	"Gin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.UserRouter(r)
	router.ClassRouter(r)
	r.Run(":8080")
}
