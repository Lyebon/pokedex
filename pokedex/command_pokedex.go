package main

import (
	"fmt"
)

func commandPokedex(cfg *config, rest ...string) error {
	fmt.Println("Your pokedex:")
	for key, _:= range cfg.caught{
		fmt.Println("-", key)
	}
	return nil
}