package services

import (
	"course/internal/domain/models"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)



//--------------Mock-------------
type MockCourseRepository struct{
	mock.Mock
}
func(m *MockCourseRepository) GetCourseById(id uint)(*models.Curso, error){
	args := m.Called(id)
	if(args.Get(0) == nil){
		return nil, args.Error(1)
	}

	return args.Get(0).(*models.Curso), args.Error(1)
}
func (m *MockCourseRepository)GetAllCourseById(id uint)(*models.Curso, error){
	return nil, nil
}
func (m *MockCourseRepository)GetModuleById(id uint)(*models.Modulo, error){
	return nil, nil
}
func (m *MockCourseRepository)SaveCourse(*models.Curso)error{
	return nil
}
func (m *MockCourseRepository)SaveModule(module *models.Modulo)error{
	return nil
}
func (m *MockCourseRepository)SaveTheme(theme *models.Tema)error{
	return nil
}
func (m *MockCourseRepository)UpdateCourse(id uint, course *models.Curso)error{
	return nil
}
func (m *MockCourseRepository)UpdateModule(id uint, module *models.Modulo)error{
	return nil
}
func (m *MockCourseRepository)UpdateTheme(id uint, theme *models.Tema)error{
	return nil
}
func (m *MockCourseRepository)DeleteCourseById(id uint)error{
	return nil
}
func (m *MockCourseRepository)DeleteModuleById(id uint)error{
	return nil
}
func (m *MockCourseRepository)DeleteThemeById(id uint)error{
	return nil
}


//---------------test---------------

func TestGetCourseById(t *testing.T){
	mockRepo := new(MockCourseRepository)
	service := NewCourseService(mockRepo)

	courseId := uint(1)

	cursoExistente := &models.Curso{
		ID: 1,
		NombreCurso: "Curso de prueba",
		NombreTutor: "Tutor de prueba",
		VideoPresentacion: "Url de prueba video presentacion",
		MetasAprendizaje: "metas de aprendizaje prueba",
	}


	mockRepo.On("GetCourseById", courseId).Return(cursoExistente, nil)


	resp, erro := service.GetCourseById(courseId)
	assert.NoError(t, erro)
	assert.NotNil(t, resp)
	mockRepo.AssertExpectations(t)

}