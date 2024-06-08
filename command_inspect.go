package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you have not caught this pokemon")
	}

	fmt.Printf("Name: %s \n", pokemon.Name)
	fmt.Printf("Height: %v \n", pokemon.Height)
	fmt.Printf("Weight: %v \n", pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("%v: %v \n", stat.Stat.Name, stat.BaseStat)
	}
	for _, typ := range pokemon.Types {
		fmt.Printf("Type: %s \n", typ.Type.Name)
	}

	return nil
}
