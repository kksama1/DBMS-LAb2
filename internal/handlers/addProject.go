package handlers

import (
	"net/http"
)

func AddProject(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../lab2/static/addProject.html")
}
