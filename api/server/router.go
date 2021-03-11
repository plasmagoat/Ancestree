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
		user := new(controllers.PersonController)
		tree := new(controllers.TreeController)
		personGroup := api.Group("person")
		{
			personGroup.GET("/:id", user.Get)
			//personGroup.GET("/parents/:id", user.GetParents)
			//personGroup.GET("/siblings/:id", user.GetSiblings)
			//personGroup.GET("/children/:id", user.GetChildren)
			personGroup.POST("/", user.Create)
			//personGroup.PUT("/:id", user.Update)
			personGroup.POST("/parents/:id", user.CreateParent)
			personGroup.POST("/child/:id", user.CreateChild)
			personGroup.POST("/link/:childId/:parentId", user.CreateLink)
			//personGroup.DELETE(":id", user.Delete)
		}
		treeGroup := api.Group("tree")
		{
			treeGroup.GET("/:id", tree.GetFamiliyTree)
		}
		// relationGroup := api.Group("profile")
		// {
		// 	profile := new(controllers.ProfileController)
		// 	relationGroup.GET("/:id", profile.GetProfile)
		// }

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
