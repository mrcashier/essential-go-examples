package pages

import (
	"github.com/russross/blackfriday"
	"html/template"
	"io"
	"io/ioutil"
	"log"
)

type Page struct {
	Title  string
	Author string
	HTML   string
	File   string
}

func NewPage(filename string) (*Page, error) {
	p := &Page{File: filename}
	return p, p.load()
}

func (p *Page) Render(layout string, w io.Writer) error {
	t, e := template.ParseFiles(layout)

	if e != nil {
		log.Fatalln(e)
	}

	return t.Execute(w, p)
}

func (p *Page) load() error {
	c, err := LoadConfig("config.json")
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(p.File)
	if err != nil {
		return err
	}

	p.HTML = string(blackfriday.MarkdownCommon(data))
	p.Title = c.Name()
	p.Author = c.Author()

	return nil
}
