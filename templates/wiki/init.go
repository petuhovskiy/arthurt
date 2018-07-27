package wiki

import "net/http"

// InitHandlers function inits `net/http` handlers for wiki pages.
func InitHandlers() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
}
