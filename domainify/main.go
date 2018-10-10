package main

import (
	"math/rand"
	"time"
	"bufio"
	"os"
	"strings"
	"unicode"
)

var tlds = []string {"com", "net"}
const allowCachars = "abcdefghijklmnopqrstuvwxyz0123456789_-"
func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan(){
		text := strings.ToLower(s.Text())
		var newText []rune
		for _,r := range text{
			if unicode.IsSpace(r){
				r = '-'
			}

			if ! strings.ContainsRune(allowCachars, r){
				continue
			}
			newText = append(newText, r)
		}

		println(string(newText)+"."+tlds[rand.Intn(len(tlds))])
	}
}
