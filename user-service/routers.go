package apiuser

import "github.com/gin-gonic/gin"

func Add(r *gin.Engine) *gin.Engine {
	route := r.Group("/users")
	{
		route.GET("/", GetUsers)
		route.POST("/", AddUser)
		route.GET("/:id", GetUser)
		route.PUT("/:id", UpdateUser)
		route.DELETE("/:id", DeleteUser)
	}
	return r
}
