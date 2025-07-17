package services

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"TorChat/internal/models"
)

// Pokemon type alias for external exposure
type Pokemon = models.Pokemon

// PokemonService handles Pokemon-related operations
type PokemonService struct {
	client *http.Client
}

// NewPokemonService creates a new Pokemon service
func NewPokemonService() *PokemonService {
	return &PokemonService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// PokeAPIResponse represents the response from PokeAPI
type PokeAPIResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
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
func (s *PokemonService) GetRandomPokemon() (*Pokemon, error) {
	// Generate random Pokemon ID up to the sinnoh region (1 - 493 for the sinnoh region)
	pokemonID := rand.Intn(493) + 1

	// Fetch Pokemon data
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", pokemonID)
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Pokemon: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var data PokeAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
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
