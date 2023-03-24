package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

	}
}
