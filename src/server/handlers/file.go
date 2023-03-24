package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	m "server/database/models"
	. "server/types"

	"gorm.io/gorm"
)

func Metadata(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	log.Println("Uploading file")

	var meta_data File

	switch r.Method {

	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Kaam 25 hain..."))

	case "POST":
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Panic(err)
		}

		err = json.Unmarshal(body, &meta_data)

		if err != nil {
			log.Panic(err)
		}

		new_file := m.File{
			Name:     meta_data.Name,
			Size:     meta_data.Size,
			Mime:     meta_data.Mime,
			IsPublic: false,
		}

		// PLEASE DO NOT FORGET VALIDATION

		result := db.Create(&new_file)

		file_data, err := json.MarshalIndent(&new_file, "", "	")

		if result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Please try again later."))
			log.Panic(err)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(file_data)
	}
}
