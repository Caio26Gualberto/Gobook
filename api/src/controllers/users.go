package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUserRepository(db)
	repository.Create(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários!"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário pelo ID!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário!"))
}
