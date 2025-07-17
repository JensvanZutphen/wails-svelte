package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// PokeAPIResponse represents the response from PokeAPI
type PokeAPIResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

// GetRandomPokemon fetches a random Pokemon from the PokeAPI
func (a *App) GetRandomPokemon() (*Pokemon, error) {
	// Generate random Pokemon ID (1-151 for original Pokemon)
	rand.Seed(time.Now().UnixNano())
	pokemonID := rand.Intn(151) + 1

	// Fetch Pokemon data
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", pokemonID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data PokeAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	// Extract types
	var types []string
	for _, t := range data.Types {
		types = append(types, t.Type.Name)
	}

	return &Pokemon{
		ID:       data.ID,
		Name:     data.Name,
		ImageURL: data.Sprites.FrontDefault,
		Types:    types,
	}, nil
}