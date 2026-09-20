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
	//api.GET("/:id", courseController.GetAllCourseById)
	api.GET("/course/:id", courseController.GetAllCourseById)
	api.POST("/course", courseController.SaveCourse)
	api.POST("/course/module/:idCourse", courseController.SaveModule)
	api.POST("/course/module/theme/:idModule", courseController.SaveTheme)
	
	
	return routers
}
