package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)


func InitRoutes() {
	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Key string `json:"key"`
			URL string `json:"url"`
		}
		if err:=json.NewDecoder(r.Body).Decode(&body); err!=nil {
			http.Error(w, "Invalid Request Body", http.StatusBadRequest)
			return
		}
		fmt.Println("got request with", body)
		if !strings.HasPrefix(body.URL, "http://") &&
				!strings.HasPrefix(body.URL, "https://") {
					body.URL = "https://" + body.URL
				}
		if err := createShortCut(body.Key, body.URL); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Shortcut created",
		})

	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		var body struct{
			Key string `json:"key"`
		}
		if err:=json.NewDecoder(r.Body).Decode(&body); err!=nil {
				http.Error(w, "Invalid Request Body", http.StatusBadRequest)
				return
		}
		if err:=deleteShortCut(body.Key); err!=nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}	
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message":"deleted successfully!",
		})
		
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key := strings.TrimPrefix(r.URL.Path, "/")
		if key == "" {
			http.Error(w, "Missing shortcut key", http.StatusBadRequest)
			return
		}
		redirect, err:=getShortcut(key)
		fmt.Println("got ", redirect, " for ", key)
		if err != nil {
			fmt.Println("error in retrieval:", err)
			http.Error(w, "Failed to retrieve shortcut", http.StatusInternalServerError)
			return
		}
		if redirect == "" {
			http.NotFound(w, r)
			return
		}

		http.Redirect(w,r,redirect, http.StatusFound)
	})
}