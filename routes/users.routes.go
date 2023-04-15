package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/lautarooyuela/recetapp-backend/db"
	"github.com/lautarooyuela/recetapp-backend/models"
	"github.com/lautarooyuela/recetapp-backend/security"
	"gorm.io/gorm"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	token := r.Header["Token"][0]

	user.Email = security.TakeEmail(token)

	if err := db.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			json.NewEncoder(w).Encode("No existing user")
		} else {
			log.Println("Error on finding user in database")
		}
	} else {
		db.DB.Where("email = ?", user.Email).Find(&user.Recipes)

		json.NewEncoder(w).Encode(&user)
	}
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	token := r.Header["Token"][0]
	email := security.TakeEmail(token)

	json.NewDecoder(r.Body).Decode(&user)
	user.Email = email

	createdUser := db.DB.Create(&user)

	if createdUser.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(createdUser.Error.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func PatchUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	token := r.Header["Token"][0]
	email := security.TakeEmail(token)

	json.NewDecoder(r.Body).Decode(&user)
	updatedUser := db.DB.Where("email = ?", email).UpdateColumns(&user)
	err := updatedUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var recipes []models.Recipe

	token := r.Header["Token"][0]
	user.Email = security.TakeEmail(token)

	if err := db.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			json.NewEncoder(w).Encode("No existing user")
		} else {
			log.Println("Error on finding user in database")
		}
	} else {
		db.DB.Where("email = ?", user.Email).Find(&recipes)

		db.DB.Unscoped().Delete(&recipes)
		db.DB.Unscoped().Delete(&user)
		w.WriteHeader(http.StatusOK)
	}

}
