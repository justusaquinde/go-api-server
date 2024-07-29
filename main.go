package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{ID: "1", Name: "Item 1"},
	{ID: "2", Name: "Item 2"},
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	items = append(items, newItem)
	json.NewEncoder(w).Encode(newItem)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	items = append(items, newItem)
	json.NewEncoder(w).Encode(newItem)
}

func main() {
	http.HandleFunc("/items", getItems)
	http.HandleFunc("/item", getItem)
	http.HandleFunc("/create", createItem)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}