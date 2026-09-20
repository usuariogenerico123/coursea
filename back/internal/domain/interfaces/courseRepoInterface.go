package interfaces
import "course/internal/domain/models"


type CourseRepositoryInterface interface{


	GetAllCourseById(id uint)(*models.Curso, error)
	GetModuleById(id uint)(*models.Modulo, error)
	GetCourseById(id uint)(*models.Curso, error)
	SaveCourse(*models.Curso)error
	SaveModule(module *models.Modulo)error
	SaveTheme(theme *models.Tema)error

}

