package main

import (
	"time"

	"github.com/vehspota/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	cfg.commands = getCommands()
	startRepl(cfg)
}
