package main

import (
	"time"

	"github.com/vehspota/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}
	cfg.commands = getCommands()
	startRepl(cfg)
}
