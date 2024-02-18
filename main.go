package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Per0na/golang-pokemon-api/models"
)

// Pokemon array to store all the pokemons
var pokemons = []models.Pokemon{
	{Id: 1, Name: "Pikachu", Type: "Electric", Level: 1},
	{Id: 2, Name: "Charmander", Type: "Fire", Level: 1},
	{Id: 3, Name: "Bulbasaur", Type: "Grass", Level: 1},
	{Id: 4, Name: "Squirtle", Type: "Water", Level: 1},
}

func main() {
	server := http.NewServeMux()

	// Find all pokemons from the pokemon array
	server.HandleFunc("/v1/pokemons", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pokemons)
	})

	// Run the server
	fmt.Println("Server running at port 8080")
	http.ListenAndServe(":8080", server)
}
