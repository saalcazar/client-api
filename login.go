package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func loginClient(url, email, password string) LoginResponse {
	login := Login{
		Email:    email,
		Password: password,
	}

	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(&login)
	if err != nil {
		log.Fatalf("error en marshal de login: %v", err)
	}

	resp := httpClient(http.MethodPost, url, "", data)
	defer resp.Body.Close()

	body, err := io.ReadAll((resp.Body))
	if err != nil {
		log.Fatalf("Error en lectura del Body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Se esperaba código 200, se obtuvo: %d, respuesta: %s", resp.StatusCode, string(body))
	}

	dataResponse := LoginResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("Error en el unmarshal del body: %v", err)
	}

	return dataResponse
}
