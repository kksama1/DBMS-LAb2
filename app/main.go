package main

import (
	"net/http"
	"projext/internal/handlers"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.MainPage)
	router.HandleFunc("/addProject", handlers.AddProject)
	//router.HandleFunc("/addNewTasks", handlers.AddNewTasks)
	router.HandleFunc("/postProject", handlers.PostProject)
	router.HandleFunc("/searchByFio", handlers.SearchByFio)
	router.HandleFunc("/searchResult", handlers.SearchResult)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
