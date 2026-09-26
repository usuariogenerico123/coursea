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
	Modulos []Modulo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
func (c *Curso)AddData(courseData course.CourseInsertDTO)*Curso{
	c.NombreCurso = courseData.NombreCurso
	c.NombreTutor = courseData.NombreTutor
	c.VideoPresentacion = courseData.VideoPresentacion
	c.MetasAprendizaje = courseData.MetasAprendizaje
	return c
}
func (c *Curso) Update(data course.CourseInsertDTO)*Curso{

	c.NombreCurso = data.NombreCurso
	c.NombreTutor = data.NombreTutor
	c.VideoPresentacion = data.VideoPresentacion
	c.MetasAprendizaje = data.MetasAprendizaje

	return c
}
func (c *Curso)verifyBlank(old string, new string)string{
	if(new == ""){
		return old
	}
	return new
}



type Modulo struct{
	gorm.Model
	TituloModulo string 
	NumeroModulo int
	DescripcionModulo string
	CursoID uint
	Curso Curso
	Temas []Tema `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
func (m *Modulo)AddData (courseId uint, moduleData course.ModuloInserDTO)*Modulo{
	m.TituloModulo = moduleData.TituloModulo
	m.NumeroModulo = moduleData.NumeroModulo
	m.DescripcionModulo = moduleData.DescripcionModulo 
	m.CursoID = courseId
	return m
}
func (m *Modulo) Update(data course.ModuloInserDTO)*Modulo{

	m.TituloModulo = data.TituloModulo
	m.NumeroModulo = data.NumeroModulo
	m.DescripcionModulo = data.DescripcionModulo

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
func (t *Tema) Update(data course.TemaInsertDTO)*Tema{

	t.NumeroTema = data.NumeroTema
	t.TituloTema = data.TituloTema
	t.UrlVideo = data.UrlVideo
	t.Duracion = data.Duracion
	t.Descriptcion = data.Descriptcion
	t.MetasAprendizaje = data.MetasAprendizaje

	return t
}

