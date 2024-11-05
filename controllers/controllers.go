package controllers

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetPesonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func FindPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personality
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var p models.Personality
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Create(&p)
	json.NewEncoder(w).Encode(p)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personality
	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personality
	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}
