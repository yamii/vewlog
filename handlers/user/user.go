package user

import(
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"

	"vewlog/models"
)

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "hiii",
	})
}

func createUser(c *gin.Context) {
	user := &models.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	str := fmt.Sprintf("name=%s, username=%s, password=%s", user.Name, user.Username, user.Password)
	fmt.Println(str)

	user.Create()
	c.JSON(http.StatusOK, gin.H{
		"hello": "Created",
	})
}

var handlermap = map[string] func(*gin.Context) {
	"GET_USER" : getUsers,
	"CREATE_USER" : createUser,
}

var valmap = map[string] govalidator.MapData {
	"GET_USER" : govalidator.MapData {

	},

	"CREATE_USER" : govalidator.MapData {
		"name" : []string{"required"},
		"username" : []string{"required"},
		"password" : []string{"required"},
	},
}

func handlerWrapper(alias string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO : validation

		// TODO: what if the function is not defined what will happen? handle it
		handlermap[alias](c)
	}
}

func SetupRoutes(router *gin.Engine) {
	// Version one
	v1 := router.Group("/api/v1")
	v1.GET("/users", handlerWrapper("GET_USER"))
	v1.POST("/users", handlerWrapper("CREATE_USER"))
}
