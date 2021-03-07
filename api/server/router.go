package server

import (
	"api/controllers"
	"api/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	InitSwaggerInfo()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.Use(middlewares.AuthMiddleware())

	api := router.Group("api")
	{
		personGroup := api.Group("person")
		{
			user := new(controllers.PersonController)
			personGroup.GET("/:id", user.Get)
			personGroup.POST("/", user.Create)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Page not found"})
	})

	return router
}

func InitSwaggerInfo() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Family Tree API"
	docs.SwaggerInfo.Description = "API for connecting to database"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080/"
	docs.SwaggerInfo.BasePath = "api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
