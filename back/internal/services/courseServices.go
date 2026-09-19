package services

import (
	"course/internal/domain/course"
	"course/internal/domain/interfaces"
	"course/internal/domain/models"
	"fmt"
)



type CourseServices struct{
	Repo interfaces.CourseRepositoryInterface
}



func NewCourseService(repo interfaces.CourseRepositoryInterface)interfaces.CourseServicesInterface{
	return &CourseServices{Repo: repo}
}

func (c *CourseServices) Service(){
	resp, erro := c.Repo.GetAllItems()
	if(erro != nil){
		fmt.Println(erro)
		return
	}
	fmt.Println("jijiijij")
	fmt.Println(resp)
}



func (c *CourseServices) SaveCourse(courseData course.CourseInsertDTO)(*course.CourseResponseDTO, error){
	var courses models.Course
	var courseResponse course.CourseResponseDTO

	courses = models.Course{}
	courses.AddData(courseData)

	resp := c.Repo.SaveCourse(&courses)
	if(resp != nil){
		return &course.CourseResponseDTO{}, resp
	}
	courseResponse = course.CourseResponseDTO{
		Id: courses.ID,
		NombreCurso: courses.NombreCurso,
		NombreTutor: courses.NombreTutor,
		VideoPresentacion: courses.VideoPresentacion,
		MetasAprendizaje: courses.MetasAprendizaje,
	}

	return &courseResponse, nil
	//fmt.Println("coursedata", courseData)
	//fmt.Println(&courses)

}

func (c *CourseServices) SaveModule(courseId uint, moduleData course.ModuloInserDTO)(*course.ModuloResponseDTO, error){
	var moduleModel *models.Modulo
	var moduleResponse course.ModuloResponseDTO

	_, err := c.Repo.GetCourseById(courseId)
	if(err != nil){
		return nil, err
	}

	moduleModel = &models.Modulo{}
	moduleModel.AddData(moduleData)

	resp := c.Repo.SaveModule(moduleModel)
	if(resp != nil){
		return nil, resp
	}
	moduleResponse = course.ModuloResponseDTO{
		Id: moduleModel.ID,
		TituloModulo: moduleModel.TituloModulo,
		NumeroModulo: moduleModel.NumeroModulo,
		DescripcionModulo: moduleModel.DescripcionModulo,

	}
	return &moduleResponse, nil
}




