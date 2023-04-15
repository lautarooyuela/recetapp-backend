package security

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lautarooyuela/recetapp-backend/models"
)

func CreateJWT(account models.Account) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["email"] = account.Email

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return tokenStr
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized"))
				}
				return []byte(os.Getenv("SECRET")), nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized: " + err.Error()))
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized token is null"))
		}
	})
}

func TakeEmail(tokenString string) string {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("no valid algorithm")
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	if !token.Valid {
		fmt.Println(err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println(err.Error())
	}

	email := claims["email"].(string)

	return email

}
