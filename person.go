package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func createPerson(url, token string, person *Person) GeneralResponse {
	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(person)
	if err != nil {
		log.Fatalf("error en marshal de persona: %v", err)
	}

	res := httpClient(http.MethodPost, url, token, data)
	defer res.Body.Close()

	body, err := io.ReadAll((res.Body))
	if err != nil {
		log.Fatalf("Error en lectura del Body: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		log.Fatalf("Se esperaba código 201, se obtuvo: %d, respuesta: %s", res.StatusCode, string(body))
	}

	dataResponse := GeneralResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("Error en el unmarshal del body: %v", err)
	}
	return dataResponse
}
