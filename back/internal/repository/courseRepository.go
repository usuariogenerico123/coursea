package repository

import (
	"course/internal/domain/interfaces"
	"course/internal/domain/models"
	"errors"

	"gorm.io/gorm"
)




type CourseRepository struct{
	Db *gorm.DB
}


func NewCourseRepository(db *gorm.DB)interfaces.CourseRepositoryInterface{
	return &CourseRepository{Db: db}
}


func (c *CourseRepository)GetCourseById(courseId uint)(*models.Course, error){
	var courseModel models.Course
	resp := c.Db.Where("id = ?", courseId).First(&courseModel)
	
	if(errors.Is(resp.Error, gorm.ErrRecordNotFound)){
		return nil, gorm.ErrRecordNotFound
	}
	if(resp.Error != nil){
		return nil, resp.Error
	}
	return &courseModel, nil

}

func(c *CourseRepository)GetAllItems()(*models.Course, error){
	var course models.Course 
	resp := c.Db.Preload("Modulos").Preload("Modulos.Temas").First(&course)
	if(resp.Error != nil){
		return nil, resp.Error
	}
	return  &course, nil
}


func (c *CourseRepository) SaveCourse(course *models.Course)error{

	resp := c.Db.Save(course)
	if (resp != nil){
		return resp.Error
	}
	return nil
}

func (c *CourseRepository) SaveModule(module *models.Modulo)error{
	resp := c.Db.Save(module)
	if (resp != nil){
		return resp.Error
	}
	
	return nil
}
func (c *CourseRepository) SaveTheme(theme *models.Tema)error{
	
	return nil
}

