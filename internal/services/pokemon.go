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
	Height  int    `json:"height"`
	Weight  int    `json:"weight"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
	} `json:"abilities"`
	Species struct {
		URL string `json:"url"`
	} `json:"species"`
}

// PokeAPISpeciesResponse represents the species response from PokeAPI
type PokeAPISpeciesResponse struct {
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
		} `json:"language"`
	} `json:"flavor_text_entries"`
	EvolutionChain struct {
		URL string `json:"url"`
	} `json:"evolution_chain"`
}

// PokeAPIEvolutionChainResponse represents the evolution chain response from PokeAPI
type PokeAPIEvolutionChainResponse struct {
	Chain EvolutionChainLink `json:"chain"`
}

// EvolutionChainLink represents a link in the evolution chain
type EvolutionChainLink struct {
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	EvolutionDetails []struct {
		MinLevel int `json:"min_level"`
		Trigger  struct {
			Name string `json:"name"`
		} `json:"trigger"`
		Item struct {
			Name string `json:"name"`
		} `json:"item"`
		HeldItem struct {
			Name string `json:"name"`
		} `json:"held_item"`
	} `json:"evolution_details"`
	EvolvesTo []EvolutionChainLink `json:"evolves_to"`
}

// PokeAPIPokemonBasicResponse represents a basic Pokemon response for evolution data
type PokeAPIPokemonBasicResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
}

// PokeAPIAbilityResponse represents the ability response from PokeAPI
type PokeAPIAbilityResponse struct {
	EffectEntries []struct {
		Effect   string `json:"effect"`
		Language struct {
			Name string `json:"name"`
		} `json:"language"`
	} `json:"effect_entries"`
}

// GetRandomPokemon fetches a random Pokemon from the PokeAPI
func (s *PokemonService) GetRandomPokemon() (*Pokemon, error) {
	// Generate random Pokemon ID up to the sinnoh region (1 - 493 for the sinnoh region)
	pokemonID := rand.Intn(493) + 1
	return s.GetPokemonByID(pokemonID)
}

// GetPokemonByID fetches a specific Pokemon by ID from the PokeAPI
func (s *PokemonService) GetPokemonByID(pokemonID int) (*Pokemon, error) {
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

	// Extract stats
	var stats []models.PokemonStat
	for _, s := range data.Stats {
		stats = append(stats, models.PokemonStat{
			Name:     s.Stat.Name,
			BaseStat: s.BaseStat,
		})
	}

	// Extract abilities with descriptions
	var abilities []models.PokemonAbility
	for _, a := range data.Abilities {
		description, err := s.fetchAbilityDescription(a.Ability.URL)
		if err != nil {
			// If we can't fetch the description, use empty string as fallback
			description = ""
		}
		abilities = append(abilities, models.PokemonAbility{
			Name:        a.Ability.Name,
			Description: description,
			IsHidden:    a.IsHidden,
		})
	}

	// Fetch species data (description and evolutions)
	species, err := s.fetchSpeciesData(data.Species.URL)
	if err != nil {
		// If we can't fetch the species data, use empty values as fallback
		species = models.PokemonSpecies{
			Description: "",
			Evolutions:  []models.PokemonEvolution{},
		}
	}

	return &Pokemon{
		ID:        data.ID,
		Name:      data.Name,
		ImageURL:  data.Sprites.FrontDefault,
		Types:     types,
		Stats:     stats,
		Abilities: abilities,
		Height:    data.Height,
		Weight:    data.Weight,
		Species:   species,
	}, nil
}

// fetchAbilityDescription fetches the description for a Pokemon ability
func (s *PokemonService) fetchAbilityDescription(url string) (string, error) {
	resp, err := s.client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch ability: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var data PokeAPIAbilityResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("failed to decode ability response: %w", err)
	}

	// Find English description
	for _, entry := range data.EffectEntries {
		if entry.Language.Name == "en" {
			return entry.Effect, nil
		}
	}

	return "", nil
}

// fetchSpeciesData fetches the species data including description and evolutions
func (s *PokemonService) fetchSpeciesData(url string) (models.PokemonSpecies, error) {
	resp, err := s.client.Get(url)
	if err != nil {
		return models.PokemonSpecies{}, fmt.Errorf("failed to fetch species: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.PokemonSpecies{}, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var data PokeAPISpeciesResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return models.PokemonSpecies{}, fmt.Errorf("failed to decode species response: %w", err)
	}

	// Find English flavor text
	var description string
	for _, entry := range data.FlavorTextEntries {
		if entry.Language.Name == "en" {
			description = entry.FlavorText
			break
		}
	}

	// Fetch evolution chain
	var evolutions []models.PokemonEvolution
	if data.EvolutionChain.URL != "" {
		evolutions, err = s.fetchEvolutionChain(data.EvolutionChain.URL)
		if err != nil {
			// If we can't fetch evolutions, continue with empty slice
			evolutions = []models.PokemonEvolution{}
		}
	}

	return models.PokemonSpecies{
		Description: description,
		Evolutions:  evolutions,
	}, nil
}

// fetchEvolutionChain fetches the evolution chain for a Pokemon
func (s *PokemonService) fetchEvolutionChain(url string) ([]models.PokemonEvolution, error) {
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch evolution chain: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var data PokeAPIEvolutionChainResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode evolution chain response: %w", err)
	}

	// Extract all Pokemon in the evolution chain
	var evolutions []models.PokemonEvolution
	s.extractEvolutions(&data.Chain, &evolutions)

	return evolutions, nil
}

// extractEvolutions recursively extracts all Pokemon from the evolution chain
func (s *PokemonService) extractEvolutions(link *EvolutionChainLink, evolutions *[]models.PokemonEvolution) {
	// Get Pokemon data for this evolution
	pokemonData, err := s.fetchBasicPokemonData(link.Species.Name)
	if err == nil {
		evolution := models.PokemonEvolution{
			Name:     pokemonData.Name,
			ID:       pokemonData.ID,
			ImageURL: pokemonData.Sprites.FrontDefault,
		}
		*evolutions = append(*evolutions, evolution)
	}

	// Process all evolutions with their requirements
	for _, evolutionLink := range link.EvolvesTo {
		// Continue recursively - this will add the evolved Pokemon and its requirements
		s.extractEvolutions(&evolutionLink, evolutions)
	}
}

// fetchBasicPokemonData fetches basic Pokemon data for evolution info
func (s *PokemonService) fetchBasicPokemonData(name string) (*PokeAPIPokemonBasicResponse, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Pokemon data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var data PokeAPIPokemonBasicResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode Pokemon response: %w", err)
	}

	return &data, nil
}
