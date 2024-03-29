package adapter

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
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

func GetPokemonResponse(number int) (PokemonResponse, error) {
	response, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d/", number))
	if err != nil {
		slog.Warn("error retrieving pokemon: " + err.Error())
		var empty PokemonResponse
		return empty, err
	}
	defer response.Body.Close()
	responseData, _ := io.ReadAll(response.Body)

	var responseObject PokemonResponse
	json.Unmarshal(responseData, &responseObject)
	return responseObject, nil
}
