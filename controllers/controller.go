package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/DedeMarantes/times/database"
	"github.com/DedeMarantes/times/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAllTimes(w http.ResponseWriter, r *http.Request) {
	var times []models.Time
	database.DB.Find(&times)
	json.NewEncoder(w).Encode(times)
}

func GetAllJogadores(w http.ResponseWriter, r *http.Request) {
	var jogadores []models.Jogador
	database.DB.Find(&jogadores)
	json.NewEncoder(w).Encode(jogadores)
}

func CriaTime(w http.ResponseWriter, r *http.Request) {
	var time models.Time
	//Converte json do corpo da requisição para a struct Time
	json.NewDecoder(r.Body).Decode(&time)
	database.DB.Create(&time)
	json.NewEncoder(w).Encode(time)
}
