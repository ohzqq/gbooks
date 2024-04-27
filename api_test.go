package main

import (
	"fmt"
	"testing"

	"github.com/ohzqq/gbooks/tui"
)

func TestApiCall(t *testing.T) {
	req := testSearchRequest()
	vols, err := Search(req.String())
	if err != nil {
		t.Error(err)
	}

	ui := tui.New(vols)
	books, err := ui.Run()
	if err != nil {
		t.Error(err)
	}

	for _, vol := range books {
		fmt.Printf("%#v\n", vol)
	}
}
