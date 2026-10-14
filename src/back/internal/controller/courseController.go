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


// GetAllCourseById godoc
// @Summary      Obtener curso por ID
// @Description  Devuelve todo el curso asociados a un ID específico
// @Tags         courses
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID del recurso"
// @Success      200  {object}  map[string]interface{}  "Lista de cursos"
// @Failure      404  {object}  map[string]string       "Error de validación o de negocio"
// @Router       /{id} [get]
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




//-----------UPDATE-----------
func (c *CourseController) UpdateCourse(h *gin.Context){

	var courseBody course.CourseInsertDTO
	
	param := h.Param("id")
	id, err := strconv.ParseUint(param, 10, 0);
	if  err != nil{
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":err.Error()})
		return
	} 

	if erro := h.BindJSON(&courseBody); erro != nil{
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":erro.Error()})
		return
	}
	
	resp, er := c.Serv.UpdateCourse(uint(id), &courseBody)
	if(er != nil){
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":er.Error()})
		return
	}

	h.IndentedJSON(http.StatusOK, gin.H{"data":resp})

}

func (c *CourseController) UpdateModule(h *gin.Context){
	var moduleBody course.ModuloInserDTO
	param := h.Param("id")
	id, err := strconv.ParseUint(param, 10, 0);
	if  err != nil{
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":err.Error()})
		return
	} 
	if erro := h.BindJSON(&moduleBody); erro != nil{
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":erro.Error()})
		return
	}
	resp, er := c.Serv.UpdateModule(uint(id), &moduleBody)
	if(er != nil){
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":er.Error()})
		return
	}

	h.IndentedJSON(http.StatusOK, gin.H{"data":resp})
}

func (c *CourseController) UpdateTheme(h *gin.Context){
	var themeBody course.TemaInsertDTO
	
	param := h.Param("id")
	id, err := strconv.ParseUint(param, 10, 0);
	if  err != nil{
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":err.Error()})
		return
	} 
	if erro := h.BindJSON(&themeBody); erro != nil{
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":erro.Error()})
		return
	}
	resp, er := c.Serv.UpdateTheme(uint(id), &themeBody)
	if(er != nil){
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":er.Error()})
		return
	}

	h.IndentedJSON(http.StatusOK, gin.H{"data":resp})
}


//--------------  DELETE  -----------
func (c *CourseController) DeleteCourse(h *gin.Context){
	param := h.Param("id")
	id,err := strconv.ParseUint(param, 10, 0)
	if(err != nil){
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":err.Error()})
		return
	}
	if resp := c.Serv.DeleteCourse(uint(id)); resp != nil{
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":resp.Error()})
		return
	}
	h.Status(http.StatusNoContent)
}

func (c *CourseController) DeleteModule(h *gin.Context){
	param := h.Param("id")
	id,err := strconv.ParseUint(param, 10, 0)
	if(err != nil){
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":err.Error()})
		return
	}
	if resp := c.Serv.DeleteModule(uint(id)); resp != nil{
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":resp.Error()})
		return
	}
	h.Status(http.StatusNoContent)
}
func (c *CourseController) DeleteTheme(h *gin.Context){
	param := h.Param("id")
	id,err := strconv.ParseUint(param, 10, 0)
	if(err != nil){
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":err.Error()})
		return
	}
	if resp := c.Serv.DeleteTheme(uint(id)); resp != nil{
		h.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data":resp.Error()})
		return
	}
	h.Status(http.StatusNoContent)
}




