package gbooks

import (
	"io"
	"os"
	"testing"

	"github.com/ohzqq/audbk"
	"github.com/ohzqq/cdb"
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

	if len(books) < 1 {
		t.Errorf("got %d, expected more", len(books))
	}

	err = WriteFFMeta(os.Stdout, books[0])
	if err != nil {
		t.Error(err)
	}

	//for _, b := range books {
	//}
}

func WriteFFMeta(w io.Writer, book cdb.Book) error {
	ff := audbk.NewFFMeta()
	audbk.BookToFFMeta(ff, book.StringMap())

	_, err := ff.WriteTo(w)
	if err != nil {
		return err
	}

	return nil
}
