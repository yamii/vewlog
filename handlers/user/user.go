package user

import(
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "hi",
	})
}

func SetupRoutes(router *gin.Engine) {
	// Version one
	v1 := router.Group("/api/v1")
	v1.GET("/users", getUsers)
}
