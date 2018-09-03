package handler

import(
	"github.com/gin-gonic/gin"

	user "vewlog/handlers/user"
)

func SetupRoutes(router *gin.Engine) {
	// Version one
	user.SetupRoutes(router)
}
