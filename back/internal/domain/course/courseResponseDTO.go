package course



type CourseResponseDTO struct{
	Id uint 					`json:"id"`
	NombreCurso string 			`json:"nombre_curso"`
	NombreTutor string 			`json:"nombre_tutor"`
	VideoPresentacion string 	`json:"video_presentacion"`
	MetasAprendizaje string		`json:"metas_aprendizaje"`
	AdminID uint				`json:"admin_id"`
	Modulos []ModuloResponseDTO `json:"modulos"`
	
}


type ModuloResponseDTO struct{
	Id uint 					`json:"id"`
	TituloModulo string 		`json:"titulo_modulo"`
	NumeroModulo int			`json:"numero_modulo"`
	DescripcionModulo string	`json:"descripcion_modulo"`
	CourseID uint				`json:"curso_id"`			
	Temas []TemaResponseDTO		`json:"temas"`
}


type TemaResponseDTO struct{
	Id uint 					`json:"id"`
	NumeroTema string 			`json:"numero_tema"`
	TituloTema string 			`json:"titulo_tema"`
	UrlVideo string 			`json:"url_vide"`
	Duracion int				`json:"duracion"`
	Descriptcion string 		`json:"descripcion"`
	MetasAprendizaje string		`json:"metas_aprendizaje"`			
	ModuloID uint				`json:"modulo_id"`

}

