package config

import (
	"course/internal/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func LoadHandlers(courseController *controller.CourseController)(*gin.Engine ){
	
	routers := gin.Default() 
	
	routers.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string)bool{return true} ,//solo para pruebas
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	
	api := routers.Group("/api")
	api.GET("", courseController.GetAll)
	api.POST("/course", courseController.SaveCourse)
	api.POST("/module/:idCourse", courseController.SaveModule)
	api.POST("/theme/:idModule", courseController.SaveTheme)
	
	
	return routers
}
