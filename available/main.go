package main

import (
	"net"
	"bufio"
	"strings"
	"os"
	"fmt"
	"log"
	"time"
)

var marks = map[bool] string{true:"Domain available", false:"Domain not available"}

func exists(domain string) (bool,error){
	const whoIsServer string = "com.whois-servers.net"
	conn, err := net.Dial("tcp", whoIsServer+":43")

	if err != nil {
		println("SERVER FAULT")
		return false, err

	}

	defer conn.Close()
	conn.Write([]byte(domain + "rn"))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		if strings.Contains(strings.ToLower(scanner.Text())	, "no match"){
			println("NO MATCH")
			return false, nil
		}
	}

	return true, nil
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan(){
		domain := s.Text()
		fmt.Print(domain, " ")
		exist, err := exists(domain)
		if err != nil {
			log.Fatalln(err)
		}
	if exist {
		println("EXISTS je true")
	}
		fmt.Println(marks[!exist])
		time.Sleep(1 * time.Second)
	}
}
