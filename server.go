package main

import(
	"github.com/gin-gonic/gin"

	handler "vewlog/handlers"
)

func main() {
	router := gin.Default()

	handler.SetupRoutes(router)

	router.Run(":8088")
}
