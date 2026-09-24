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




func (c *CourseRepository)GetAllCourseById(id uint)(*models.Curso, error){
	var courseModel models.Curso
	resp := c.Db.Preload("Modulos").Preload("Modulos.Temas").First(&courseModel, id)
	if(resp.Error != nil){
		return nil, resp.Error
	}
	return &courseModel, nil
}


func (c *CourseRepository)GetModuleById(id uint)(*models.Modulo, error){
	var moduleModel models.Modulo
	resp := c.Db.Where("id = ?", id).First(&moduleModel)
	if(errors.Is(resp.Error, gorm.ErrRecordNotFound)){
		return nil, errors.New("Module not found")
	}

	if(resp.Error != nil){
		return nil, resp.Error
	}
	return &moduleModel, nil
}


func (c *CourseRepository)GetCourseById(id uint)(*models.Curso, error){
	var courseModel models.Curso
	resp := c.Db.Where("id = ?", id).First(&courseModel)
	
	if(errors.Is(resp.Error, gorm.ErrRecordNotFound)){
		return nil, gorm.ErrRecordNotFound
	}
	if(resp.Error != nil){
		return nil, resp.Error
	}
	return &courseModel, nil

}




func (c *CourseRepository) SaveCourse(course *models.Curso)error{

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
	
	resp := c.Db.Save(theme)
	if(errors.Is(resp.Error, gorm.ErrRecordNotFound)){
		return errors.New("Module not found")
	}

	if(resp.Error != nil){
		return resp.Error
	}
	return nil
}

func (c *CourseRepository) UpdateCourse(id uint, course *models.Curso)error{
	resp := c.Db.Model(&models.Curso{}).Where("id = ?", id).Updates(course)
	if(resp.Error != nil){
		return resp.Error
	}
	return nil
}
func (c *CourseRepository) UpdateModule(id uint, module *models.Modulo)error{
	resp := c.Db.Model(&models.Modulo{}).Where("id = ?", id).Updates(module)
	if(resp.Error != nil){
		return resp.Error
	}
	return nil
}
func (c *CourseRepository) UpdateTheme(id uint, theme *models.Tema)error{
	resp := c.Db.Model(&models.Tema{}).Where("id = ?", id).Updates(theme)
	if(resp.Error != nil){
		return resp.Error
	}
	return nil
}


func(c *CourseRepository)DeleteCourseById(id uint)error{
	var courseModel models.Curso
	resp := c.Db.Where("id = ?", id).First(&courseModel)
	if(resp.Error != nil){
		return resp.Error
	}
	if del := c.Db.Delete(&courseModel); del.Error != nil {
		return del.Error
	}
	return nil

}

func (c *CourseRepository) DeleteModuleById(id uint)error {
	var moduleModel models.Modulo
	if module := c.Db.Where("id = ?", id).First(&moduleModel); module.Error != nil{
		return module.Error
	}

	if del := c.Db.Delete(&moduleModel); del.Error != nil{
		return del.Error
	}
	return nil


}

func (c *CourseRepository)DeleteThemeById(id uint)error{
	var themeModel models.Tema
	if resp := c.Db.Where("id = ?", id).First(themeModel); resp.Error != nil{
		return resp.Error
	}
	if del := c.Db.Delete(&themeModel); del.Error != nil{
		return del.Error
	}
	return nil
}