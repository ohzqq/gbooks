package main

import (
	"encoding/json"
	"net/http"

	gb "google.golang.org/api/books/v1"
)

var Client = &Api{
	client: &http.Client{},
}

type Api struct {
	client *http.Client
}

func Search(q string) (gb.Volumes, error) {
	var vols gb.Volumes

	res, err := Client.client.Get(q)
	if err != nil {
		return vols, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&vols)
	if err != nil {
		return vols, err
	}

	return vols, nil
}
