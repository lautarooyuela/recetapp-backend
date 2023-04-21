package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lautarooyuela/recetapp-backend/db"
	"github.com/lautarooyuela/recetapp-backend/models"
	"github.com/lautarooyuela/recetapp-backend/routes"
	"github.com/lautarooyuela/recetapp-backend/security"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// database connection
	db.DBConnection()
	// db.DB.Migrator().DropTable(models.User{})
	db.DB.AutoMigrate(models.Account{})
	db.DB.AutoMigrate(models.Recipe{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	// Index route

	s := r.PathPrefix("/api").Subrouter()

	s.HandleFunc("/healthy", routes.Healthy).Methods("GET")

	s.HandleFunc("/login/{token}", routes.Login).Methods("GET")
	s.HandleFunc("/register", routes.Register).Methods("POST")

	// users routes
	s.Handle("/user", security.ValidateJWT(routes.GetUserHandler)).Methods("GET")
	s.Handle("/user", security.ValidateJWT(routes.PostUserHandler)).Methods("POST")
	s.Handle("/user", security.ValidateJWT(routes.PatchUserHandler)).Methods("PATCH")
	s.Handle("/user", security.ValidateJWT(routes.DeleteUserHandler)).Methods("DELETE")

	// recipes routes
	s.Handle("/recipes", security.ValidateJWT(routes.GetRecipesHandler)).Methods("GET")
	s.Handle("/recipes/{id}", security.ValidateJWT(routes.GetRecipeHandler)).Methods("GET")
	s.Handle("/recipes/{id}", security.ValidateJWT(routes.PatchRecipeHandler)).Methods("PATCH")
	s.Handle("/recipes", security.ValidateJWT(routes.CreateRecipeHandler)).Methods("POST")
	s.Handle("/recipes/{id}", security.ValidateJWT(routes.DeleteRecipeHandler)).Methods("DELETE")

	log.Println("Conexión en puerto" + os.Getenv("PORT"))
	http.ListenAndServe("0.0.0.0"+os.Getenv("PORT"), r)

}
