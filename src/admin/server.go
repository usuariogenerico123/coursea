package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("static/")))

	server := &http.Server{
		Addr: ":3009",
		Handler: mux,
	}

	fmt.Println("Servidor iniciado")

	if resp := server.ListenAndServe(); resp != nil{
		fmt.Println(resp)
	} 



}
