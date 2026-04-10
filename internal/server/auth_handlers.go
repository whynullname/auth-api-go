package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	userJsonData := getUserJsonData(r)
	if userJsonData == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func AuthUser(w http.ResponseWriter, r *http.Request) {
	userJsonData := getUserJsonData(r)
	if userJsonData == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func getUserJsonData(r *http.Request) *userDataInput {
	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "application/json") {
		return nil
	}

	userRegistationJson := userDataInput{}
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Printf("Can't read request body %v\n", err)
		return nil
	}

	err = json.Unmarshal(body, &userRegistationJson)
	if err != nil {
		log.Printf("Can't unmarshal body %v\n", err)
		return nil
	}

	return &userRegistationJson
}
