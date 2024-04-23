package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Planet struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Gravity        string   `json:"gravity"`
	Residents      []string `json:"residents"`
}

type PlanetsResponse struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Planet `json:"results"`
}

func main() {
	// membaca 1 data =========

	// response, err := http.Get("http://swapi.py4e.com/api/planets/1")

	// if err != nil {
	// 	log.Fatal("error consume API URL")
	// }

	// responseData, errBody := io.ReadAll(response.Body)
	// if errBody != nil {
	// 	log.Fatal("error real response body")
	// }
	// defer response.Body.Close()

	// var planet1 Planet
	// json.Unmarshal(responseData, &planet1)

	// fmt.Println("---- Print Field ----")
	// fmt.Println(planet1.Name)
	// fmt.Println(planet1.RotationPeriod)
	// fmt.Println(planet1.OrbitalPeriod)
	// fmt.Println(planet1.Gravity)
	// // fmt.Println(planet1.Residents)
	// for i := 0; i < len(planet1.Residents); i++ {
	// 	fmt.Println(planet1.Residents[i])
	// }

	// membaca banyak data =========
	response, err := http.Get("http://swapi.py4e.com/api/planets")

	if err != nil {
		log.Fatal("error consume API URL")
	}

	responseData, errBody := io.ReadAll(response.Body)
	if errBody != nil {
		log.Fatal("error real response body")
	}
	defer response.Body.Close()

	var allPlanets PlanetsResponse
	json.Unmarshal(responseData, &allPlanets)

	fmt.Println("---- Print Field ----")
	fmt.Println(allPlanets.Count)
	fmt.Println(*allPlanets.Next)
	fmt.Println(allPlanets.Previous)

	// fmt.Println(allPlanets.Results)
	for _, v := range allPlanets.Results {
		fmt.Println(v.Name)
		if len(v.Residents) > 0 {
			for _, penghuni := range v.Residents {
				fmt.Println(penghuni)
			}
		}
	}

}
