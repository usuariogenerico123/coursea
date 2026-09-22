package interfaces

import (
	"course/internal/domain/course"
	"course/internal/domain/models"
)


type CourseServicesInterface interface{
	
	GetCourseById(id uint)(*models.Curso, error)
	GetAllCourseById(id uint)(*course.CursoResponseDTO, error)
	SaveCourse(courseData course.CourseInsertDTO)(*course.CursoResponseDTO, error)
	SaveModule(courseId uint, moduleData course.ModuloInserDTO)(*course.ModuloResponseDTO, error)
	SaveTheme(moduleId uint, themeData course.TemaInsertDTO)(*course.TemaResponseDTO, error)
	UpdateCourse(courseId uint, dataCourse *course.CourseInsertDTO)(*course.CursoResponseDTO, error)
	UpdateModule(moduleId uint, dataModule *course.ModuloInserDTO)(*course.ModuloResponseDTO, error)
	UpdateTheme(themeId uint, dataTheme *course.TemaInsertDTO)(*course.TemaResponseDTO, error)
	

}	

