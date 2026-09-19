package course





type CourseInsertDTO struct{
	
	NombreCurso string 			`json:"nombre_curso"`
	NombreTutor string 			`json:"nombre_tutor"`
	VideoPresentacion string 	`json:"video_presentacion"`
	MetasAprendizaje string		`json:"metas_aprendizaje"`
	
	Modulos []ModuloInserDTO 	`json:"modulos"`
	
}


type ModuloInserDTO struct{

	TituloModulo string 		`json:"titulo_modulo"`
	NumeroModulo int			`json:"numero_modulo"`
	DescripcionModulo string	`json:"descripcion_modulo"`
	CourseID uint				`json:"course_id"`
	Temas []TemaInsertDTO		`json:"temas"`
}


type TemaInsertDTO struct{

	NumeroTema string 			`json:"numero_tema"`
	TituloTema string 			`json:"titulo_tema"`
	UrlVideo string 			`json:"url_vide"`
	Duracion int				`json:"duracion"`
	Descriptcion string 		`json:"descripcion"`
	MetasAprendizaje string		`json:"metas_aprendizaje"`				
	ModuloID uint 				`json:"modulo_id"`

}

