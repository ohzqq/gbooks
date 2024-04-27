//go:build exclude

package main

import (
	"strings"

	"github.com/ohzqq/ur/pkg/book"
	"golang.org/x/exp/slices"
	gb "google.golang.org/api/books/v1"
)

type Meta struct {
	*gb.Volume
}

func (m Meta) Get(key string) string {
	switch key {
	case book.Authors:
		return strings.Join(m.VolumeInfo.Authors, " & ")
	case book.Cover:
		if c := m.VolumeInfo.ImageLinks.Medium; c != "" {
			return c
		}
	case book.Description:
		if d := m.VolumeInfo.Description; d != "" {
			return d
		}
	case book.Identifiers:
		var id []string
		for _, i := range m.VolumeInfo.IndustryIdentifiers {
			id = append(id, i.Type+":"+i.Identifier)
		}
		return strings.Join(id, ",")
	case book.Languages:
		if l := m.VolumeInfo.Language; l != "" {
			return l
		}
	case book.Publisher:
		if p := m.VolumeInfo.Publisher; p != "" {
			return p
		}
	case book.Published:
		if pd := m.VolumeInfo.PublishedDate; pd != "" {
			return pd
		}
	case book.Title:
		return m.VolumeInfo.Title
	}
	return ""
}

func (m Meta) Has(key string) bool {
	return slices.Contains(m.Keys(), key)
}

func (m Meta) Keys() []string {
	return []string{
		"authors",
		"comments",
		"identifiers",
		"languages",
		"publisher",
		"pubdate",
		"title",
	}
}
