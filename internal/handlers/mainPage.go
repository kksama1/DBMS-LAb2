package handlers

import "net/http"

func MainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../lab2/static/index.html")
}
