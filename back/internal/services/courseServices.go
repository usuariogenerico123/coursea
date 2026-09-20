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

func (c *CourseServices)GetAllCourseById(id uint)(*course.CursoResponseDTO, error){
	var modules []course.ModuloResponseDTO
	//var themes []course.TemaResponseDTO

	resp, erro := c.Repo.GetAllCourseById(id)
	if(erro != nil){
		fmt.Println(erro)
		return nil, erro
	}


	for _,v := range(resp.Modulos){
		temas := []*course.TemaResponseDTO{}
		for _, x := range(v.Temas){
			temas = append(temas, &course.TemaResponseDTO{
				Id: x.ID,
				NumeroTema: x.NumeroTema,
				TituloTema: x.TituloTema,
				UrlVideo: x.UrlVideo,
				Duracion: x.Duracion,
				Descriptcion: x.Descriptcion,
				MetasAprendizaje: x.MetasAprendizaje,
				ModuloID: x.ModuloID,
			})
		}
		modules = append(modules, course.ModuloResponseDTO{
			Id: v.ID,
			TituloModulo: v.TituloModulo,
			NumeroModulo: v.NumeroModulo,
			DescripcionModulo: v.DescripcionModulo,
			CursoID: v.CursoID,
			Temas: temas,
		})
	}


	courseReponse := course.CursoResponseDTO{
		Id: resp.ID,
		NombreCurso: resp.NombreCurso,
		NombreTutor: resp.NombreTutor,
		VideoPresentacion: resp.VideoPresentacion,
		MetasAprendizaje: resp.MetasAprendizaje,
		Modulos: modules,
	}



	return &courseReponse, nil
}


//----------------- GET     ---------------
func (c *CourseServices)GetCourseById(id uint)(*models.Curso, error){

	resp, erro := c.Repo.GetCourseById(id)
	if(erro != nil){
		fmt.Println(erro)
		return nil, erro
	}
	fmt.Println(resp)
	return resp, nil

}



func (c *CourseServices) SaveCourse(courseData course.CourseInsertDTO)(*course.CursoResponseDTO, error){
	var courses models.Curso
	var courseResponse course.CursoResponseDTO

	courses = models.Curso{}
	courses.AddData(courseData)

	resp := c.Repo.SaveCourse(&courses)
	if(resp != nil){
		return &course.CursoResponseDTO{}, resp
	}
	courseResponse = course.CursoResponseDTO{
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
	moduleModel.AddData(courseId, moduleData)

	resp := c.Repo.SaveModule(moduleModel)
	if(resp != nil){
		return nil, resp
	}
	moduleResponse = course.ModuloResponseDTO{
		Id: moduleModel.ID,
		TituloModulo: moduleModel.TituloModulo,
		NumeroModulo: moduleModel.NumeroModulo,
		DescripcionModulo: moduleModel.DescripcionModulo,
		CursoID: courseId,
	}
	return &moduleResponse, nil
}


func (c *CourseServices)SaveTheme(moduleId uint, themeData course.TemaInsertDTO)(*course.TemaResponseDTO, error){
	theme := models.Tema{}
	theme.AddTema(moduleId, themeData)

	_, er := c.Repo.GetModuleById(moduleId)
	if(er != nil){
		
		return nil, er
	}


	resp := c.Repo.SaveTheme(&theme)
	if(resp != nil){
		return nil, resp
	}

	themeResponse := course.TemaResponseDTO{
		Id: theme.ID,
		NumeroTema: theme.NumeroTema,
		TituloTema: theme.TituloTema,
		UrlVideo: theme.UrlVideo,
		Duracion: theme.Duracion,
		Descriptcion: theme.Descriptcion,
		MetasAprendizaje: theme.MetasAprendizaje,
		ModuloID: moduleId,
	}

	return &themeResponse, nil
}










