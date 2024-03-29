package adapter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PokemonResponse struct {
	Id      int            `json:"id"`
	Name    string         `json:"name"`
	Sprites PokemonSprites `json:"sprites"`
	Species PokemonSpecies `json:"species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonSprites struct {
	BackImage    string       `json:"back_default"`
	FrontImage   string       `json:"front_default"`
	OtherSprites OtherSprites `json:"other"`
}

type OtherSprites struct {
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
}

type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
}

func GetPokemonResponse(number int) PokemonResponse {
	response, _ := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d/", number))
	responseData, _ := ioutil.ReadAll(response.Body)
	var responseObject PokemonResponse
	json.Unmarshal(responseData, &responseObject)
	return responseObject
}
