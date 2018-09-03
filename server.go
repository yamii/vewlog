package main

import(
	"os"
	"log"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"

	handler "vewlog/handlers"
)

func init() {
	// Load ENV file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := gin.Default()

	// Serve Static Files
	router.Use(static.Serve("/", static.LocalFile("./views/build", true)))

	handler.SetupRoutes(router)

	port := os.Getenv("PORT")
	router.Run(":" + port)
}
