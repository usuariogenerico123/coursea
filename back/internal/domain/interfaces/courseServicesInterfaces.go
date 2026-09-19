package interfaces

import "course/internal/domain/course"


type CourseServicesInterface interface{

	Service()
	SaveCourse(courseData course.CourseInsertDTO)(*course.CourseResponseDTO, error)
	SaveModule(courseId uint, moduleData course.ModuloInserDTO)(*course.ModuloResponseDTO, error)
}	

