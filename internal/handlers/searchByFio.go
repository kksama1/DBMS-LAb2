package handlers

import (
	"net/http"
)

func SearchByFio(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../lab2/static/searchByFio.html")
}
