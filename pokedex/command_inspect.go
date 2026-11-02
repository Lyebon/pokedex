package main

import (
	"fmt"
)

func commandInspect(cfg *config, rest ...string) error {
	poke := rest[0]
	pokedata, ok := cfg.caught[poke]
	if !ok{
		fmt.Printf("%s was not catch", poke)
		return nil
	}
	fmt.Printf(`
Name: %v
Height: %v
Weight: %v
Stats:
`, pokedata.Name, pokedata.Height, pokedata.Weight)
	for _, stat := range pokedata.Stats{
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, poketype := range pokedata.Types{
		fmt.Printf("  -%s\n", poketype.Type.Name)
	}
	return nil
}