package wiki

import "io/ioutil"

type page struct {
	Title string
	Body  []byte
}

func getPageFilename(title string) string {
	return title + ".txt"
}

func (p *page) save() error {
	filename := getPageFilename(p.Title)
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*page, error) {
	filename := getPageFilename(title)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &page{Title: title, Body: body}, nil
}
