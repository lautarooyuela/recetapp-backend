package routes

import (
	"log"
	"net/http"
)

func Healthy(w http.ResponseWriter, r *http.Request) {
	log.Println("Healthy")
	w.Write([]byte("Healthy"))
}
