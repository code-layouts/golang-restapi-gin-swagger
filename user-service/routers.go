package apiuser

import "github.com/gin-gonic/gin"

func Add(r *gin.Engine) *gin.Engine {
	r.RouterGroup.GET("/users", GetUsers)
	r.RouterGroup.GET("/users/:id", GetUser)
	r.RouterGroup.POST("/users", AddUser)
	r.RouterGroup.PUT("/users/:id", UpdateUser)
	r.RouterGroup.DELETE("/users/:id", DeleteUser)
	return r
}
