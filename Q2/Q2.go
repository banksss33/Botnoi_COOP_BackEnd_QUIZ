package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PokemonDetail struct {
	Stats []struct {
		Base_stat int `json:"base_stat"`
		Effort    int `json:"effort"`
		Stat      struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Name    string `json:"name"`
	Sprites struct {
		Back_default       string `json:"back_default"`
		Back_female        string `json:"back_female"`
		Back_shiny         string `json:"back_shiny"`
		Back_shiny_female  string `json:"back_shiny_female"`
		Front_default      string `json:"front_default"`
		Front_female       string `json:"front_female"`
		Front_shiny        string `json:"front_shiny"`
		Front_shiny_female string `json:"front_shiny_female"`
	} `json:"sprites"`
}

func main() {
	router := gin.Default()
	router.POST("/pokemon-detail", retrievePokemonDetail)
	router.Run(":8080")
}

func retrievePokemonDetail(c *gin.Context) {
	pokemonID := 1

	pokemonDetail := fetchPokemonData(pokemonID)
	c.JSON(http.StatusOK, pokemonDetail)
}

func fetchPokemonData(pokemonID int) PokemonDetail {
	pokemonEndpoint := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d/", pokemonID)
	pokemonFormEndpoint := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-form/%d/", pokemonID)

	pokemonDetail := PokemonDetail{}

	pokemonResponse, err := http.Get(pokemonEndpoint)
	if err != nil {
		fmt.Println("Error fetching data from pokemon endpoint:", err)
		return pokemonDetail
	}
	defer pokemonResponse.Body.Close()

	pokemonBody, err := io.ReadAll(pokemonResponse.Body)
	if err != nil {
		fmt.Println("Error reading response from pokemon endpoint:", err)
		return pokemonDetail
	}

	err = json.Unmarshal(pokemonBody, &pokemonDetail)
	if err != nil {
		fmt.Println("Error decoding JSON from pokemon endpoint:", err)
		return pokemonDetail
	}

	pokemonFormResponse, err := http.Get(pokemonFormEndpoint)
	if err != nil {
		fmt.Println("Error fetching data from pokemon-form endpoint:", err)
		return pokemonDetail
	}
	defer pokemonFormResponse.Body.Close()

	pokemonFormBody, err := io.ReadAll(pokemonFormResponse.Body)
	if err != nil {
		fmt.Println("Error reading response from pokemon-form endpoint:", err)
		return pokemonDetail
	}

	var pokemonFormData PokemonDetail
	err = json.Unmarshal(pokemonFormBody, &pokemonFormData)
	if err != nil {
		fmt.Println("Error decoding JSON from pokemon-form endpoint:", err)
		return pokemonDetail
	}

	if pokemonFormData.Name != "" {
		pokemonDetail.Name = pokemonFormData.Name
	}
	if pokemonFormData.Sprites.Front_default != "" {
		pokemonDetail.Sprites.Front_default = pokemonFormData.Sprites.Front_default
	}

	return pokemonDetail
}
