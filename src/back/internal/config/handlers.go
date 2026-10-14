package config

import (
	"course/internal/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
	
)


func LoadHandlers(courseController *controller.CourseController)(*gin.Engine ){
	
	routers := gin.Default() 
	
	routers.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string)bool{return true} ,//solo para pruebas
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	
	api := routers.Group("/api/course")

	//api.GET("/:id", courseController.GetAllCourseById)
	api.GET("/:id", courseController.GetAllCourseById)
	//-------POST----------
	api.POST("/", courseController.SaveCourse)
	api.POST("/module/:idCourse", courseController.SaveModule)
	api.POST("/module/theme/:idModule", courseController.SaveTheme)
	//-------UPDATE-----
	api.PUT("/:id", courseController.UpdateCourse)
	api.PUT("/module/:id", courseController.UpdateModule)
	api.PUT("/module/theme/:id", courseController.UpdateTheme)
	//-------DELETE-----
	api.DELETE("/:id", courseController.DeleteCourse)
	api.DELETE("/module/:id", courseController.DeleteModule)
	api.DELETE("/module/theme/:id", courseController.DeleteTheme)
	
	//-------DOC---------
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return routers
}
