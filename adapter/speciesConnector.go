package adapter

import (
	"encoding/json"
	"io/ioutil"
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
	responseData, _ := ioutil.ReadAll(response.Body)
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
