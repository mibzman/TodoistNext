package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func NotionPostReq(url string, body io.Reader) error {

	// Create a Bearer string by appending string access token
	var bearer = "Bearer <secret>"

	// Create a new request using http
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", "2021-05-13")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func AddToToTable(TableID string, TitleText string) error {
	var jsonData = []byte(fmt.Sprintf(`{
		"parent": { "database_id": "%v" },
		"properties": {
		  "Name": {
			"title": [
			  {
				"text": {
				  "content": "%v"
				}
			  }
			]
		  }
		}
	  }`, TableID, TitleText))

	return NotionPostReq("https://api.notion.com/v1/pages", bytes.NewBuffer(jsonData))
}
