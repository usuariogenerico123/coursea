package repository

import (
	"course/internal/domain/models"
	
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)



func SetUpDb(t *testing.T)*gorm.DB{
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Curso{})
	assert.NoError(t, err)

	return db

}


func TestDeleteCourseById(t *testing.T){
	db := SetUpDb(t)
	repo := NewCourseRepository(db)
	course := &models.Curso{NombreCurso: "curso de prueba"}

	err := db.Create(&course).Error
	assert.NoError(t, err)
	assert.NotZero(t, course.ID)


	err = repo.DeleteCourseById(course.ID)
	assert.NoError(t, err)

	var foundCourse models.Curso
	err = db.First(&foundCourse, course.ID).Error
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

}




