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
	router.Use(gin.Recovery())

	router.Static("/resources", "./swagger/api/resources")

	// swagger UI: url pointing to API definition
	// http://localhost:8080/swagger/index.html
	routeRoot := router.Group("/")
	{
		routeRoot.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "server page") })
		routeRoot.GET("/openapi.yaml",
			func(c *gin.Context) {
				if gin.Mode() == gin.ReleaseMode {
					f.Println("gin.ReleaseMode", gin.ReleaseMode)
					c.String(404, "ReleaseMode 에선 API 를 지원하지 않습니다. openapi.yaml not found")
					return
				}
			},
			func(c *gin.Context) { c.File("./swagger/api/openapi.yaml") },
		)

		swaggerConfig := &ginSwagger.Config{URL: "/openapi.yaml"}
		routeRoot.GET("/swagger/*any", ginSwagger.CustomWrapHandler(swaggerConfig, swaggerFiles.Handler))
	}
	return router
}

func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
	f.Println("GetHealth")
}

func main() {
	router := Router()
	router.GET("/health", GetHealth)
	apiuser.Add(router)
	router.Run(":8080")
}
