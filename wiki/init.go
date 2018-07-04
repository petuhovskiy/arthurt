package wiki

import "net/http"

// Attach wiki handlers to `net/http`
func InitHandlers() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
}
