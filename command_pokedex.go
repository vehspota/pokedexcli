package main

import (
	"fmt"
	"sort"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty. Go catch some Pokemon!")
		return nil
	}

	names := make([]string, 0, len(cfg.caughtPokemon))
	for name := range cfg.caughtPokemon {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Println("Your Pokedex:")
	for _, name := range names {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
