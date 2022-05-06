package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func TodoistGetReq(url string) []byte {

	// Create a Bearer string by appending string access token
	var bearer = "Bearer <secret>"

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	return body
}

func TodoistPostReq(url string, body io.Reader) {

	// Create a Bearer string by appending string access token
	var bearer = "Bearer <secret>"

	// Create a new request using http
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	req.Header.Add("Content-Type", "application/json")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()
}
