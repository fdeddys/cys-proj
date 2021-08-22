package router

import (
	"github.com/fdeddys/tes/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           86400,
	}))

	UserController := new(controller.UserController)
	api := r.Group("/api")
	{
		api.POST("/user/register", UserController.SaveDataUser)
		api.POST("/user/gather", UserController.GatherData)
		api.POST("/user/list-resources", UserController.ListResources)
	}

	return r

}
