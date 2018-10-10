package main

import (
	"net/http"
	"errors"
	"encoding/json"
	"bufio"
	"os"
	"log"
	"fmt"
)

const(
	requestURL string = "http://words.bighugelabs.com/api/2/fef37bf83d046f297a05a7519ca86150/"

)

type Synonym struct {
	Noun noun `json:"noun"`
}

type  noun struct {
	Synonyms []string `json:"syn"`
}

func Synonyms(word string) (Synonym, error){
	response, err := http.Get(requestURL+word+"/json")

	if err != nil {
		return Synonym{}, errors.New("Failed to find synonyms for word ["+word+"], err: "+err.Error())
	}
	var data Synonym
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return Synonym{}, err
	}
	return data, nil
}

func main() {

	s := bufio.NewScanner(os.Stdin)
	for s.Scan(){
		word := s.Text()
		syns, err := Synonyms(word)
		if err != nil{
			log.Fatalln("Error in searching for synonyms. "+err.Error())
		}

		if len(syns.Noun.Synonyms) == 0 {
			log.Fatalln("Could not find any synonyms for given word: "+word)
		}
		fmt.Println(word)
		for _, syn := range syns.Noun.Synonyms {
			fmt.Println(syn)
		}
	}
}
