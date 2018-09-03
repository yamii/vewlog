package main

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"

	handler "vewlog/handlers"
)

func main() {
	router := gin.Default()

	handler.SetupRoutes(router)
	// Serve Static Files
	router.Use(static.Serve("/", static.LocalFile("./views/build", true)))

	router.Run(":8088")
}
