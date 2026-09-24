package interfaces
import "course/internal/domain/models"


type CourseRepositoryInterface interface{


	GetAllCourseById(id uint)(*models.Curso, error)
	GetModuleById(id uint)(*models.Modulo, error)
	GetCourseById(id uint)(*models.Curso, error)
	SaveCourse(*models.Curso)error
	SaveModule(module *models.Modulo)error
	SaveTheme(theme *models.Tema)error
	UpdateCourse(id uint, course *models.Curso)error
	UpdateModule(id uint, module *models.Modulo)error
	UpdateTheme(id uint, theme *models.Tema)error
	DeleteCourseById(id uint)error
	DeleteModuleById(id uint)error
	DeleteThemeById(id uint)error
	
}

