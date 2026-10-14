package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)






func LoadDb()*gorm.DB{
	

	dsnExample := "host=aws-0-us-west-2.pooler.supabase.com user=postgres.lftdvguedwvmciuyajvl password=Manchoso_123 dbname=postgres port=5432 "
	// dsn := os.Getenv("URL_DB")
	// if(len(dsn) == 0){
	// 	panic("No se pudo cargar variable de entorno URL_DB")
	// }
	DB, err := gorm.Open(postgres.Open(dsnExample), nil)
	if(err != nil){
		panic(err)
	}
	
	log.Println("Connection success")
	return DB


}

