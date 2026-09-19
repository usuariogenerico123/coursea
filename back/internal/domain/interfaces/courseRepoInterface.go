package interfaces
import "course/internal/domain/models"


type CourseRepositoryInterface interface{

	GetCourseById(courseId uint)(*models.Course, error)
	GetAllItems()(*models.Course, error)
	SaveCourse(*models.Course)error
	SaveModule(module *models.Modulo)error
	SaveTheme(theme *models.Tema)error

}

