package models

import (
	"course/internal/domain/course"

	"gorm.io/gorm"
)


type Course struct{
	gorm.Model
	NombreCurso string 
	NombreTutor string 
	VideoPresentacion string 
	MetasAprendizaje string
	Modulos []Modulo
	Admin Admin 
	AdminID uint
	
}
func (c *Course)AddData(courseData course.CourseInsertDTO)*Course{
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
	CourseID uint
	Course Course
	Temas []Tema
}
func (m *Modulo)AddData (moduleData course.ModuloInserDTO)*Modulo{
	m.TituloModulo = moduleData.TituloModulo
	m.NumeroModulo = moduleData.NumeroModulo
	m.DescripcionModulo = moduleData.DescripcionModulo 
	m.CourseID = moduleData.CourseID
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
func (t *Tema)AddTema (dataTheme course.TemaInsertDTO)*Tema{
	t.NumeroTema = dataTheme.NumeroTema
	t.TituloTema = dataTheme.TituloTema
	t.UrlVideo = dataTheme.UrlVideo
	t.Duracion = dataTheme.Duracion
	t.Descriptcion = dataTheme.Descriptcion
	t.MetasAprendizaje = dataTheme.MetasAprendizaje
	t.ModuloID = dataTheme.ModuloID
	return t
}

