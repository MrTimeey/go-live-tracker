package adapter

import (
	"encoding/json"
	"io"
	"log/slog"
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

func GetSpeciesResponse(pokemonResponse PokemonResponse) (SpeciesResponse, error) {
	response, err := http.Get(pokemonResponse.Species.Url)
	if err != nil {
		slog.Warn("error retrieving species: " + err.Error())
		var empty SpeciesResponse
		return empty, err
	}
	defer response.Body.Close()

	responseData, _ := io.ReadAll(response.Body)
	var responseObject SpeciesResponse
	json.Unmarshal(responseData, &responseObject)
	return responseObject, nil
}

func GetGermanName(response SpeciesResponse) string {
	for _, name := range response.Names {
		if name.Language.Name == "de" {
			return name.Name
		}
	}
	return response.Name
}
