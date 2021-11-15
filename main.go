package main

import (
	apiuser "example/apiserver/v1/user-service"
	f "fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(
		gin.Recovery(),
	)

	// endpoints
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "server page")
	})

	route := router.Group("/")
	{
		// swagger UI: url pointing to API definition
		// http://localhost:8080/swagger/index.html
		route.GET("/openapi.yaml",
			func(c *gin.Context) {
				if gin.Mode() == gin.ReleaseMode {
					c.String(404, "openapi.yaml not found")
					return
				}
			},
			func(c *gin.Context) { c.File("./swagger/api/openapi.yaml") },
		)

		route.GET("/swagger/*any",
			ginSwagger.WrapHandler(swaggerFiles.Handler,
				ginSwagger.URL("/openapi.yaml"),
			),
		)

	}

	return router
}

// GetHealth - health
func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	f.Println("GetHealth")
}

func main() {
	router := Router()
	router.GET("/health", GetHealth)
	apiuser.Add(router)
	router.Run(":8080")
}
