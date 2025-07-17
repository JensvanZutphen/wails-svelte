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

// PokemonEvolution represents a Pokemon evolution
type PokemonEvolution struct {
	Name         string `json:"name"`
	ID           int    `json:"id"`
	ImageURL     string `json:"image_url"`
	MinLevel     int    `json:"min_level,omitempty"`
	TriggerName  string `json:"trigger_name,omitempty"`
	ItemName     string `json:"item_name,omitempty"`
	HeldItemName string `json:"held_item_name,omitempty"`
}

// PokemonSpecies represents Pokemon species information
type PokemonSpecies struct {
	Description string             `json:"description"`
	Evolutions  []PokemonEvolution `json:"evolutions"`
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
