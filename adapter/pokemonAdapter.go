package adapter

import (
	"fmt"
	"math/rand"
)

type Pokemon struct {
	Id    int
	Name  string
	Image string
}

func GetRandomPokemon() Pokemon {
	pokemonResponse, _ := GetPokemonResponse(getRandomInt())
	speciesResponse, _ := GetSpeciesResponse(pokemonResponse)
	name := GetGermanName(speciesResponse)
	fmt.Println(speciesResponse)
	return Pokemon{
		Id:    pokemonResponse.Id,
		Name:  name,
		Image: pokemonResponse.Sprites.OtherSprites.OfficialArtwork.FrontDefault,
	}
}

func getRandomInt() int {
	min := 1
	max := 251
	return rand.Intn(max-min) + min
}
