package models

import (
	"course/internal/domain/course"
	"gorm.io/gorm"
)


type Curso struct{
	gorm.Model
	NombreCurso string 
	NombreTutor string 
	VideoPresentacion string 
	MetasAprendizaje string
	Modulos []Modulo
}
func (c *Curso)AddData(courseData course.CourseInsertDTO)*Curso{
	c.NombreCurso = courseData.NombreCurso
	c.NombreTutor = courseData.NombreTutor
	c.VideoPresentacion = courseData.VideoPresentacion
	c.MetasAprendizaje = courseData.MetasAprendizaje
	return c
}



type Modulo struct{
	gorm.Model
	TituloModulo string 
	NumeroModulo int
	DescripcionModulo string
	CursoID uint
	Curso Curso
	Temas []Tema
}
func (m *Modulo)AddData (courseId uint, moduleData course.ModuloInserDTO)*Modulo{
	m.TituloModulo = moduleData.TituloModulo
	m.NumeroModulo = moduleData.NumeroModulo
	m.DescripcionModulo = moduleData.DescripcionModulo 
	m.CursoID = courseId
	return m
}



type Tema struct{

	gorm.Model
	NumeroTema string 
	TituloTema string 
	UrlVideo string 
	Duracion int
	Descriptcion string 
	MetasAprendizaje string
	Modulo Modulo
	ModuloID uint

}
func (t *Tema)AddTema (moduloId uint, dataTheme course.TemaInsertDTO)*Tema{
	t.NumeroTema = dataTheme.NumeroTema
	t.TituloTema = dataTheme.TituloTema
	t.UrlVideo = dataTheme.UrlVideo
	t.Duracion = dataTheme.Duracion
	t.Descriptcion = dataTheme.Descriptcion
	t.MetasAprendizaje = dataTheme.MetasAprendizaje
	t.ModuloID = moduloId
	return t
}

