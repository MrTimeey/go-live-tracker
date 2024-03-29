package adapter

import (
	"encoding/json"
	"io"
	"net/http"
)

type SpeciesResponse struct {
	Name  string        `json:"name"`
	Names []PokemonName `json:"names"`
}

type PokemonName struct {
	Language LanguageStruct `json:"language"`
	Name     string         `json:"name"`
}

type LanguageStruct struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GetSpeciesResponse(pokemonResponse PokemonResponse) SpeciesResponse {
	response, _ := http.Get(pokemonResponse.Species.Url)
	responseData, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	var responseObject SpeciesResponse
	json.Unmarshal(responseData, &responseObject)
	return responseObject
}

func GetGermanName(response SpeciesResponse) string {
	for _, name := range response.Names {
		if name.Language.Name == "de" {
			return name.Name
		}
	}
	return response.Name
}
