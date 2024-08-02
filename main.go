package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Song represents a song with id, title, and artist
type Song struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

var songs []Song

func main() {
	// Initialize songs slice with some sample data
	songs = append(songs, Song{ID: 1, Title: "Song 1", Artist: "Artist 1"})
	songs = append(songs, Song{ID: 2, Title: "Song 2", Artist: "Artist 2"})
	songs = append(songs, Song{ID: 3, Title: "Song 3", Artist: "Artist 3"})

	router := mux.NewRouter()

	router.HandleFunc("/songs", getSongs).Methods("GET")
	router.HandleFunc("/songs/{id}", getSong).Methods("GET")
	router.HandleFunc("/songs", createSong).Methods("POST")
	router.HandleFunc("/songs/{id}", updateSong).Methods("PUT")
	router.HandleFunc("/songs/{id}", deleteSong).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// getSongs returns all songs
func getSongs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(songs)
}

// getSong returns a single song by ID
func getSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	for _, item := range songs {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Song not found", http.StatusNotFound)
}

// createSong creates a new book
func createSong(w http.ResponseWriter, r *http.Request) {
	var newSong Song
	err := json.NewDecoder(r.Body).Decode(&newSong)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newSong.ID = len(songs) + 1
	songs = append(songs, newSong)

	json.NewEncoder(w).Encode(newSong)
	w.WriteHeader(http.StatusCreated)
}

// updatedSong updates an existing song by ID
func updateSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	for index, item := range songs {
		if item.ID == id {
			var updatedSong Song
			err = json.NewDecoder(r.Body).Decode(&updatedSong)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			updatedSong.ID = id // Ensure ID remains consistent

			songs[index] = updatedSong

			json.NewEncoder(w).Encode(updatedSong)
			return
		}
	}
	http.Error(w, "Song not found", http.StatusNotFound)
}

// deleteSong deletes a song by ID
func deleteSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	for index, item := range songs {
		if item.ID == id {
			songs = append(songs[:index], songs[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Song not found", http.StatusNotFound)
}
