package main

// Pokemon represents a Pokemon with basic info
type Pokemon struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	ImageURL string   `json:"image_url"`
	Types    []string `json:"types"`
}