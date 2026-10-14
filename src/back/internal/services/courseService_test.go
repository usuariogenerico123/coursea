package services

import (
	"course/internal/domain/course"
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
	args := m.Called(id)
	if( args.Get(0) == nil){
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Curso), args.Error(1)

}
func (m *MockCourseRepository) GetThemeById(id uint)(*models.Tema, error){
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
	args := m.Called(id)
	return args.Error(0)
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

func TestGetAllCourseById(t *testing.T){
	mockRepo := new(MockCourseRepository)
	service := NewCourseService(mockRepo)

	id := uint(1)

	temas := []models.Tema{
		{
			ID: 1,
			NumeroTema: "1",
			TituloTema: "Titulo de prueba",
			UrlVideo: "url de prueba",
			Duracion: "2:00 h",
			Descriptcion: "Descripcion de prueba",
			MetasAprendizaje: "Metas de aprendizaje de prueba",
			ModuloID: 1,
		},
	}
	modulos := []models.Modulo{
		{
			ID: 1,
			TituloModulo: "Modulo de prueba",
			NumeroModulo: 1,
			DescripcionModulo: "Descripcion de prueba",
			CursoID: 1,
			Temas: temas,
		},
	}
	curso := models.Curso{
		ID: 1,
		NombreCurso: "Curso de pruuueba",
		NombreTutor: "Tutor de pruuueba",
		VideoPresentacion: "Url de prueba.com",
		MetasAprendizaje: "Metas de aprendizaje",
		Modulos: modulos,
	}



	respDto := &course.CursoResponseDTO{
		Id: curso.ID,
		NombreCurso: curso.NombreCurso,
		NombreTutor: curso.NombreTutor,
		VideoPresentacion: curso.VideoPresentacion,
		MetasAprendizaje: curso.VideoPresentacion,
		Modulos: []course.ModuloResponseDTO{
			{
				Id: modulos[0].ID,
				TituloModulo: modulos[0].TituloModulo,
				NumeroModulo: modulos[0].NumeroModulo,
				DescripcionModulo: modulos[0].DescripcionModulo,
				CursoID: modulos[0].CursoID,
				Temas: []*course.TemaResponseDTO{
					{
						Id: temas[0].ID,
						TituloTema: temas[0].TituloTema,
						UrlVideo: temas[0].UrlVideo,
						Duracion: temas[0].Duracion,
						Descriptcion: temas[0].Descriptcion,
						MetasAprendizaje: temas[0].MetasAprendizaje,
						ModuloID: temas[0].ModuloID,
					},
				},
			},
		},
	}
	mockRepo.On("GetAllCourseById", id).Return(&curso, nil)

	resp, err := service.GetAllCourseById(id)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.Id, respDto.Id)
	assert.Equal(t, resp.Modulos[0].Id, respDto.Modulos[0].Id)
	assert.Equal(t, resp.NombreCurso, respDto.NombreCurso)
	mockRepo.AssertExpectations(t)

}





