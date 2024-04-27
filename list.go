package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ohzqq/bubbles/list"
	"github.com/ohzqq/cdb"
	"golang.org/x/term"
)

type List struct {
	list.Model
	books []cdb.Book
}

type Item struct {
	cdb.Book
}

func NewUI(books []cdb.Book) List {
	var items []list.Item
	for _, book := range books {
		items = append(items, NewItem(book))
	}

	w, h := TermSize()

	m := list.New(items, list.NewDefaultDelegate(), w, h)
	m.SetNoLimit()

	return List{
		Model: m,
		books: books,
	}
}

func (l List) Run() ([]cdb.Book, error) {
	p := tea.NewProgram(l)
	_, err := p.Run()
	if err != nil {
		return nil, err
	}

	sel := l.ToggledItems()
	books := make([]cdb.Book, len(sel))
	for _, idx := range sel {
		books[idx] = l.books[idx]
	}

	return books
}

func (l List) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return l, tea.Quit
		}
	}

	var cmd tea.Cmd
	l.Model, cmd = l.Model.Update(msg)
	return l, cmd
}

func NewItem(p cdb.Book) Item {
	return Item{
		Book: p,
	}
}

func (i Item) authors() string {
	var auths []string
	for _, a := range i.Authors {
		auths = append(auths, a)
	}
	return strings.Join(auths, " & ")
}

func (i Item) FilterValue() string {
	return fmt.Sprintf("%s %s", i.Title, i.authors())
}

func (i Item) Title() string {
	return fmt.Sprintf("%s by %s", i.Title, i.authors())
}

func (i Item) Description() string {
	return i.Comments
}

func TermSize() (int, int) {
	w, h, _ := term.GetSize(int(os.Stdin.Fd()))
	//println(w)
	//println(h)
	return w, h
}
