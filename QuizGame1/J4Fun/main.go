package main

import (
	"flag"
	"fmt"
	"os"
	"encoding/json"
)

func main() {

	const token string = "&appid=0b520b45e3bb6ba695cd4d3da4adf80b"
	const BASE_URL string = "api.openweathermap.org/data/2.5/weather?q="

	
	city := flag.String("city","Prague", "Name of the city you want to receive the actual weather")
	flag.Parse()

	url := BASE_URL+*city+token
	fmt.Printf("URL=%s\n",url)
	var result Result
	jsonFile, err := os.Open("test-weather.json")

	defer jsonFile.Close()
	if err != nil{
		fmt.Printf("Cannot open json file, %s\n",err)
		os.Exit(1)
	}
	fmt.Println("File opened")

	jsonParser := json.NewDecoder(jsonFile)
	jsonParser.Decode(&result)

	fmt.Println(result.Main.Temperature)



}

type Result struct {
	Weather struct{
		Main string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct{
		Temperature float32 `json:"temp"`
		Pressure int `json:"pressure"`
		Humidity int `json:"humidity"`
		MinTemp float32 `json:"temp_min"`
		MaxTemp float32 `json:"temp_max"`
	} `json:"main"`
	Wind struct{
		Speed int `json:"speed"`
	} `json:"wind"`
	Sys struct{
		Country string `json:"country"`
		Sunrise int `json:"sunrise"`
		Sunset int `json:"sunset"`
	} `json:"sys"`
}