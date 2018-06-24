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

	result.printResult()



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
		Speed float32 `json:"speed"`
	} `json:"wind"`
	Clouds struct{
		all int `json:"all"`
	} `json:"clouds"`
	Sys struct{
		Country string `json:"country"`
		Sunrise string `json:"sunrise"`
		Sunset string `json:"sunset"`
	} `json:"sys"`
	Name string `json:"name"`
}

func (result Result) printResult(){

	fmt.Println("+-----------------------------------------+")
	fmt.Printf("|    Weather in location %s,%s 	       |\n",result.Name,result.Sys.Country)

	fmt.Println("+--------------------+--------------------+")

	fmt.Printf("|    Temperature     |      %.1f °C       |\n",(result.Main.Temperature - 272.15))
	fmt.Printf("|    Humidity        |      %d %s          |\n",result.Main.Humidity,"%")
	fmt.Printf("|    Pressure        |      %d %s      |\n",result.Main.Pressure,"hPa")
	fmt.Printf("|    Wind            |      %.1f %s       |\n",result.Wind.Speed,"m/s")
	fmt.Printf("|    Cloudiness      |      %d %s           |\n",result.Clouds.all,"%")
	fmt.Printf("+--------------------+--------------------+\n")


	}