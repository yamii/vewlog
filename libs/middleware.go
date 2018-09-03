package libs

import(
	"github.com/gin-gonic/gin"
)

/*********************************
 * Custom Middlewares
*********************************/
func BeforeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}

func AfterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
