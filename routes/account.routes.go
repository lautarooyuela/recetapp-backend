package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lautarooyuela/recetapp-backend/db"
	"github.com/lautarooyuela/recetapp-backend/models"
	"github.com/lautarooyuela/recetapp-backend/security"
	"github.com/lautarooyuela/recetapp-backend/services"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	encodedString := mux.Vars(r)
	account = services.Decode64(account, encodedString)

	if err := db.DB.Where("email = ? AND password = ?", account.Email, account.Password).First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			json.NewEncoder(w).Encode("No existing account")
		} else {
			log.Println("Error on finding account in database")
		}
	} else {
		var jwt = security.CreateJWT(account)
		account.Token = jwt

		json.NewEncoder(w).Encode(&account)
	}

}

func Register(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	json.NewDecoder(r.Body).Decode(&account)
	createdAccount := db.DB.Create(&account)

	if createdAccount.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(createdAccount.Error.Error()))
	}

	var user models.User
	user.Email = account.Email
	user.Name = "Nuevo Usuario"

	createdUser := db.DB.Create(&user)

	if createdUser.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(createdUser.Error.Error()))
	}

	json.NewEncoder(w).Encode("Account created")
}
