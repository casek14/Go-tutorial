package main

import (
	"net/http"
	"fmt"
	"html"
	"log"
	"flag"
	"io/ioutil"
	"encoding/json"
)

const PORT string = ":8080"
var valuesMap = map[string]string{
	"dog": "REDIRECTED TO SOME SHIT ABOUT DOGS",
	"cat": "I DO NOT LIKE CATS !!!",
	"go": "GoLang is the best programing language in the world !!!",
}
func main() {

	yamlFile := flag.String("yaml","settings.yml","Specify the yaml file to be used for redirect " +
		"handling")
	flag.Parse()

	yamlF, err := ioutil.ReadFile("settings.json")
	if err != nil{
		log.Fatal("CANNOT OPEN GIVEN FILE %v", err)
	}

	var config Configuration
	err = json.Unmarshal([]byte(yamlF),&config)


	fmt.Println(config.Redirects)
	if err != nil{
		log.Fatal("CANNOT READ GIVEN FILE %s", yamlFile)
	}


	for _,c := range config.Redirects{
		fmt.Printf("Request: %s - Redirect: %s\n",c.RequestSymbol,c.RedirectPage)
	}

	//http.HandleFunc("/", returnHome)
	http.HandleFunc("/",handleRequest)
	log.Fatal(http.ListenAndServe(PORT, nil))
}


func returnHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"<h1> Hello, you are at url: <font color=\"orange\">",html.EscapeString(r.URL.Path),"</font></h1>")
}

func handleRequest(w http.ResponseWriter, r *http.Request){
	if r.URL.Path == "/" {
		fmt.Fprint(w,"<h1>Wellcome enjoy this great Golang App made by Casek!</h1>")
	return
	}else {
		var req string = r.URL.Path[len("/"):]
		fmt.Fprint(w,"<h1> Requested URL is: <font color=\"orange\">", req,"</font></h1>")

		if value, key := valuesMap[req]; key{
			text := "Redirecting to URL= "+value
			fmt.Fprint(w,"<h1> Requested URL is: <font color=\"green\">", text,"</font></h1>")

		}
	}

}


type Configuration struct{
	Redirects []struct{
		RequestSymbol string `json:"request_symbol"`
		RedirectPage string `json:"redirect_page"`
	} `json:"redirects"`
}


