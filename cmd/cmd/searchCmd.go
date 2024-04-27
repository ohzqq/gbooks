package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gobuffalo/flect"
	"github.com/ohzqq/audbk"
	"github.com/ohzqq/cdb"
	"github.com/ohzqq/gbooks"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func searchCmdRun(cmd *cobra.Command, args []string) {
	req := gbooks.NewRequest().Keywords(args[0])
	if a := viper.GetString("author"); a != "" {
		req.Author(a)
	}
	if t := viper.GetString("title"); t != "" {
		req.Title(t)
	}
	fmt.Printf("search url %s\n", req.String())

	books, err := gbooks.Search(req.String())
	if err != nil {
		log.Fatalf("search error: %v\n", err)
	}

	switch len(books) {
	case 0:
		println("no results")
		os.Exit(0)
	case 1:
		err = processResults(books)
		if err != nil {
			log.Fatalf("results error: %v\n", err)
		}
	default:
		ui := gbooks.NewUI(books)
		sel, err := ui.Run()
		if err != nil {
			log.Fatalf("ui error: %v\n", err)
		}
		err = processResults(sel)
		if err != nil {
			log.Fatalf("results error: %v\n", err)
		}
	}
}

func processResults(books []cdb.Book) error {
	for _, book := range books {
		exts := viper.GetStringSlice("ext")
		for _, ext := range exts {
			var enc encoder
			if viper.GetBool("dont-save") {
				enc = getEnc(os.Stdout, ext)
			} else {
				name := flect.Underscore(book.Title)
				f, err := os.Create(name + ext)
				if err != nil {
					return err
				}
				defer f.Close()
				enc = getEnc(f, ext)
			}

			err := enc.Encode(book)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type encoder interface {
	Encode(v any) error
}

type ffmeta struct {
	w io.Writer
}

func (f ffmeta) Encode(v any) error {
	book := v.(cdb.Book)

	ff := audbk.NewFFMeta()
	audbk.BookToFFMeta(ff, book.StringMap())

	_, err := ff.WriteTo(f.w)
	if err != nil {
		return err
	}
	return nil
}

func getEnc(w io.Writer, ext string) encoder {
	switch ext {
	case ".json":
		return json.NewEncoder(w)
	case ".toml":
		return toml.NewEncoder(w)
	case ".ini":
		return writeFFMeta(w)
	default:
		enc := yaml.NewEncoder(w)
		enc.SetIndent(2)
		return enc
	}
}

func writeFFMeta(w io.Writer) encoder {
	return ffmeta{
		w: w,
	}
}
