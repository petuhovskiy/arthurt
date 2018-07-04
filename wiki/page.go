package wiki

import "io/ioutil"

type Page struct {
	Title string
	Body  []byte
}

func getPageFilename(title string) string {
	return title + ".txt"
}

func (p *Page) save() error {
	filename := getPageFilename(p.Title)
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := getPageFilename(title)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
