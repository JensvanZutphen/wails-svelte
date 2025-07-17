package app

import (
	"context"
	"fmt"

	"TorChat/internal/services"
)

// App struct
type App struct {
	ctx            context.Context
	pokemonService *services.PokemonService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		pokemonService: services.NewPokemonService(),
	}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetRandomPokemon fetches a random Pokemon from the PokeAPI
func (a *App) GetRandomPokemon() (*services.Pokemon, error) {
	return a.pokemonService.GetRandomPokemon()
}

// GetPokemonByID fetches a specific Pokemon by ID from the PokeAPI
func (a *App) GetPokemonByID(id int) (*services.Pokemon, error) {
	return a.pokemonService.GetPokemonByID(id)
}
