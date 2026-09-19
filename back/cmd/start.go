package main

import (
	"course/internal/config"
	"course/internal/controller"
	"course/internal/domain/models"
	"course/internal/repository"
	"course/internal/services"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB
func Start(){
	
	port := ":3006"
	DB = config.LoadDb()
	DB.AutoMigrate(&models.Course{}, &models.Modulo{}, &models.Tema{})

	//--repository--
	courseRepo := repository.NewCourseRepository(DB)
	//--services---
	courseService := services.NewCourseService(courseRepo)
	//--controllers---
	courseController := controller.NewCourseController(courseService)

	handlers := config.LoadHandlers(courseController)

	server := &http.Server{
		Addr: port,
		Handler: handlers,
		ReadHeaderTimeout: 2 * time.Second,
	}

	err := server.ListenAndServe()
	if( err != nil){
		fmt.Println(err)
		return
	}

	fmt.Println("Server iniciado")

	
}