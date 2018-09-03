package main

import(
	"os"
	"log"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"

	libs "vewlog/libs"
	handler "vewlog/handlers"
)

/*
 * TODO:
 * - support loading env files in diff environments
 *   - development, staging, production
 * - production grade logger
 * - production deployment and setup
*/
func init() {
	// Load ENV file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	libs.InitDB()
}

// Entrypoint of the Application
func main() {
	router := gin.Default()

	router.Use(libs.BeforeMiddleware())
	// Serve Static Files
	router.Use(static.Serve("/", static.LocalFile("./views/build", true)))

	handler.SetupRoutes(router)

	// after middleware - TODO: doesn't work, will revisit
	router.Use(libs.AfterMiddleware())

	port := os.Getenv("PORT")
	router.Run(":" + port)
}
