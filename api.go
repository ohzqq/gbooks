package gbooks

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ohzqq/cdb"
	gb "google.golang.org/api/books/v1"
)

var Client = &Api{
	client: &http.Client{},
}

type Api struct {
	client *http.Client
}

func Search(q string) ([]cdb.Book, error) {
	res, err := Client.client.Get(q)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var vols gb.Volumes
	err = json.NewDecoder(res.Body).Decode(&vols)
	if err != nil {
		return nil, err
	}

	var info []cdb.Book
	for _, vol := range vols.Items {
		info = append(info, volToBook(vol))
	}

	return info, nil
}

func volToBook(vol *gb.Volume) cdb.Book {
	book := cdb.Book{
		EditableFields: cdb.EditableFields{
			Title:     vol.VolumeInfo.Title,
			Publisher: vol.VolumeInfo.Publisher,
			Authors:   vol.VolumeInfo.Authors,
			Tags:      vol.VolumeInfo.Categories,
			Languages: []string{vol.VolumeInfo.Language},
			Comments:  vol.VolumeInfo.Description,
		},
	}
	date, err := time.Parse(time.DateOnly, vol.VolumeInfo.PublishedDate)
	if err != nil {
		date = time.Now()
	}
	book.Pubdate = date
	return book
}
