package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func GetData() string {
	response, _ := http.Get(fmt.Sprintf("http://pokeapi.co/api/v2/pokemon/%d/", getRandomInt()))
	responseData, _ := ioutil.ReadAll(response.Body)
	return string(responseData)
}

func getRandomInt() int {
	min := 1
	max := 151
	return rand.Intn(max-min) + min
}
