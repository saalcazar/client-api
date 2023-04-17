package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const url = "http://localhost:8080"

func main() {
	lc := loginClient(url+"/v1/login", "info@saalcazar.org", "123456")
	fmt.Println(lc)

	person := Person{
		Name:        "Samuel Alcazar",
		Age:         38,
		Communities: []Community{{"saalcazar"}, {"Golang"}},
	}
	gr := createPerson(url+"/v1/persons", lc.Data.Token, &person)
	fmt.Println(gr)
}

func httpClient(method, url, token string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("Request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request: %v", err)
	}
	return response
}
