package models

// PokemonStat represents a Pokemon's base stat
type PokemonStat struct {
	Name     string `json:"name"`
	BaseStat int    `json:"base_stat"`
}

// PokemonAbility represents a Pokemon's ability
type PokemonAbility struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsHidden    bool   `json:"is_hidden"`
}

// PokemonSpecies represents Pokemon species information
type PokemonSpecies struct {
	Description string `json:"description"`
}

// Pokemon represents a Pokemon with detailed info
type Pokemon struct {
	ID        int              `json:"id"`
	Name      string           `json:"name"`
	ImageURL  string           `json:"image_url"`
	Types     []string         `json:"types"`
	Stats     []PokemonStat    `json:"stats"`
	Abilities []PokemonAbility `json:"abilities"`
	Height    int              `json:"height"` // in decimeters
	Weight    int              `json:"weight"` // in hectograms
	Species   PokemonSpecies   `json:"species"`
}
