# Pokedex CLI

An interactive command-line Pokedex written in Go. It runs as a REPL: you type commands, and the program fetches data from the [PokeAPI](https://pokeapi.co/) over HTTP and prints the results in your terminal.

Built as part of the [Boot.dev](https://boot.dev) backend curriculum to practice working with HTTP clients, JSON parsing, and structuring a Go program.

## Features

- Interactive REPL with a command prompt
- Browse location areas of the Pokémon world, 20 at a time, forwards and backwards
- Explore a location to see which Pokémon live there
- Catch Pokémon and keep them in your own Pokedex
- Inspect a caught Pokémon's height, weight, stats, and types
- Real data fetched from the PokeAPI

## Getting started

**Requirements:** Go 1.27.2 or later.

```bash
git clone https://github.com/Lyebon/pokedex.git
cd pokedex
go run .
```

## Commands

| Command   | Description                                                      |
|-----------|------------------------------------------------------------------|
| `help`    | Display a help message with all the commands                     |
| `map`     | Display the next 20 location areas                               |
| `mapb`    | Display the previous 20 location areas                           |
| `explore` | Explore a location and list all the Pokémon available there      |
| `catch`   | Attempt to catch the selected Pokémon                            |
| `inspect` | Display the name, height, weight, stats, and type(s) of a Pokémon |
| `pokedex` | Display the list of Pokémon you have captured                    |
| `exit`    | Exit the Pokedex                                                 |

## Example

```
Pokedex > help

Welcome to the Pokedex!
Usage:
mapb: Display the previous 20 location areas
pokedex: Display the list of pokemon captured
catch: Attempt to catch the pokemon selected
exit: Exit the Pokedex
explore: Explore the map location you say and list all the pokemon available
help: Display a help message with all the commands
inspect: Display the name, height, weight, stats and type(s) of the Pokemon
map: Display the next 20 location areas
```

## What I learned

- Making HTTP requests and decoding JSON in Go
- Designing a REPL: reading input, parsing commands, dispatching to handlers
- Structuring code with structs and methods