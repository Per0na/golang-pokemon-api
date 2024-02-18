package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	server.HandleFunc("GET /v1/pokemons", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pokemons)
	})

	// Find a pokemon by id
	server.HandleFunc("GET /v1/pokemons/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		// Convert the id to int
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}

		// Find the pokemon by id
		var pokemon models.Pokemon
		for _, p := range pokemons {
			if p.Id == id {
				pokemon = p
				break
			}
		}

		// If the pokemon is not found
		if pokemon.Id == 0 {
			http.Error(w, "Pokemon not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pokemon)
	})

	// Create a new pokemon
	server.HandleFunc("POST /v1/pokemons", func(w http.ResponseWriter, r *http.Request) {
		var pokemon models.Pokemon
		err := json.NewDecoder(r.Body).Decode(&pokemon)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Set the id of the new pokemon
		pokemon.Id = len(pokemons) + 1

		// Add the new pokemon to the array
		pokemons = append(pokemons, pokemon)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(pokemon)
	})

	// Update a pokemon by id
	server.HandleFunc("PUT /v1/pokemons/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		// Convert the id to int
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}

		var pokemon models.Pokemon
		err = json.NewDecoder(r.Body).Decode(&pokemon)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Find the pokemon by id
		for i, p := range pokemons {
			if p.Id == id {
				// Update the pokemon
				pokemon.Id = id
				pokemons[i] = pokemon
				break
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pokemon)
	})

	// Run the server
	fmt.Println("Server running at port 8080")
	http.ListenAndServe(":8080", server)
}
