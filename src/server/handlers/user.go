package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	m "server/database/models"
	"server/types"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func SignUp(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	log.Println("Signing you up")

	var profile_data types.User

	switch r.Method {

	case "GET":
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - GET Method not allowed"))

	case "POST":
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Panic(err)
		}

		err = json.Unmarshal(body, &profile_data)

		if err != nil {
			log.Panic(err)
		}

		v := validator.New()

		err = v.Struct(profile_data)

		new_user := m.User{
			FirstName: profile_data.FirstName,
			LastName:  profile_data.LastName,
			Email:     profile_data.Email,
			Pin:       profile_data.Pin,
		}

		if err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				log.Println(e)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("400 - Data validation failed."))
				return
			}
		}

		result := db.Create(&new_user)

		user_data, err := json.MarshalIndent(&new_user, "", "	")

		if result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Please try again later."))
			log.Panic(err)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(user_data)
	}

	// err := json.NewDecoder(r.Body).Decode(&profile_data)

	// if err != nil {
	// 	log.Println(err)
	// 	log.Println("#")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// log.Println("=======================")
	// log.Println(profile_data)

	// expirationTime := time.Now().Add(5 * time.Minute)

	// claims := &Claims{
	// 	Username: creds.Username,
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		ExpiresAt: jwt.NewNumericDate(expirationTime),
	// 	},
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	// if err != nil {
	// 	log.Println(err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   tokenString,
	// 	Expires: expirationTime,
	// })
}
