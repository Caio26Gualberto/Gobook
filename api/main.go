package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.InitializeEnvironments()
	fmt.Printf("Rodando na porta %d\n", config.Port)
	fmt.Printf("A Connection string é %s\n", config.ConnectionString)

	fmt.Println("Rodando API!")

	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
