package controller

import (
	"course/internal/domain/course"
	"course/internal/domain/interfaces"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type CourseController struct{
	Serv interfaces.CourseServicesInterface
}

func NewCourseController (serv interfaces.CourseServicesInterface)*CourseController{
	return &CourseController{Serv: serv}
}



//-----GET-----
func (c *CourseController )GetAllCourseById(h *gin.Context){
	id := h.Param("id")
	num, err := strconv.ParseUint(id, 10, 0)
	if(err != nil){
		fmt.Println(err)
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":err.Error()})
		return
	}

	resp, er := c.Serv.GetAllCourseById(uint(num))
	if(er != nil){
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":er.Error()})
		return
	}
	h.IndentedJSON(http.StatusOK, gin.H{"data":resp})
	
}


//----------GET------------
func (c *CourseController)GetCourseById(h *gin.Context){

	//id := h.Param("id")
	id, err := strconv.ParseUint(h.Param("id"), 10, 0)
	if(err != nil){
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":err.Error()})
		return
	}

	resp, er := c.Serv.GetCourseById(uint(id))
	if(er != nil){
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":er.Error()})
		return
	}

	h.IndentedJSON(http.StatusOK, gin.H{"data":resp})

}



//----post
func (c *CourseController)SaveCourse(h *gin.Context){
	var courseData course.CourseInsertDTO 
	data := h.BindJSON(&courseData)
	if(data != nil){
		fmt.Println(data)
		
	}

	resp , err := c.Serv.SaveCourse(courseData)
	if(err!=nil){
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":err.Error()})
		return
	}

	h.IndentedJSON(http.StatusCreated, gin.H{"data":resp})
}

//---POST----
func (c *CourseController)SaveModule(h *gin.Context){
	var moduleInsertData course.ModuloInserDTO

	idCourse := h.Param("idCourse")

	er := h.BindJSON(&moduleInsertData)
	if(er != nil){
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":er})
		return
	}

	num, err := strconv.ParseUint(idCourse, 10, 0)
	if(err != nil){
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":err})
		return
	}
	resp, errr := c.Serv.SaveModule(uint(num), moduleInsertData)
	if(errr == gorm.ErrRecordNotFound){
		h.AbortWithStatusJSON(http.StatusNotFound, gin.H{"data":errr})
		return
	}
	if(errr != nil){
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":errr})
		return
	}
	h.IndentedJSON(http.StatusCreated, gin.H{"data":resp})


}
//-----post
func (c *CourseController)SaveTheme(h *gin.Context){
	var themeInserData course.TemaInsertDTO

	idModule := h.Param("idModule")
	id, er := strconv.ParseUint(idModule, 10, 0)
	if(er != nil){
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":er})
		return
	}

	body := h.BindJSON(&themeInserData)
	if(body != nil){
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":body})
		return
	}

	resp, erro := c.Serv.SaveTheme(uint(id), themeInserData)
	if(erro != nil){
		fmt.Println(erro)
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":erro.Error()})
		return
	}

	h.IndentedJSON(http.StatusCreated, gin.H{"data":resp})
}



