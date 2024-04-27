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
	for _, vol := range vols.Items {
		fmt.Printf("%#v\n", vol.VolumeInfo)
	}
}
