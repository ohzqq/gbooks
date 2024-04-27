package main

import (
	"fmt"
	"testing"
)

func TestApiCall(t *testing.T) {
	req := testSearchRequest()
	vols, err := Search(req.String())
	if err != nil {
		t.Error(err)
	}

	ui := NewUI(vols)
	books, err := ui.Run()
	if err != nil {
		t.Error(err)
	}

	for _, vol := range books {
		fmt.Printf("%#v\n", vol)
	}
}
