package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lautarooyuela/recetapp-backend/db"
	"github.com/lautarooyuela/recetapp-backend/models"
	"github.com/lautarooyuela/recetapp-backend/security"
)

func GetRecipesHandler(w http.ResponseWriter, r *http.Request) {
	var recipe []models.Recipe
	token := r.Header["Token"][0]

	email := security.TakeEmail(token)

	db.DB.Where("email = ?", email).Find(&recipe)
	json.NewEncoder(w).Encode(&recipe)
}

func GetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var recipe models.Recipe
	token := r.Header["Token"][0]

	email := security.TakeEmail(token)
	db.DB.Where("email = ?", email).First(&recipe, params["id"])

	if recipe.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Recipe not found"))
		return
	}

	json.NewEncoder(w).Encode(&recipe)
}

func CreateRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	token := r.Header["Token"][0]

	email := security.TakeEmail(token)
	json.NewDecoder(r.Body).Decode(&recipe)
	recipe.Email = email

	createdRecipe := db.DB.Create(&recipe)
	err := createdRecipe.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&recipe)
}

func PatchRecipeHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var recipe models.Recipe
	token := r.Header["Token"][0]

	email := security.TakeEmail(token)
	db.DB.Where("email = ?", email).First(&recipe, params["id"])

	if recipe.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Recipe not found"))
		return
	}

	json.NewEncoder(w).Encode(&recipe)
}

func DeleteRecipeHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var recipe models.Recipe
	token := r.Header["Token"][0]

	email := security.TakeEmail(token)
	db.DB.Where("email = ?", email).First(&recipe, params["id"])

	if recipe.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Recipe not found"))
		return
	}

	db.DB.Unscoped().Delete(&recipe)
	w.WriteHeader(http.StatusOK)
}
