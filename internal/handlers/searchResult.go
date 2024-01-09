package handlers

import (
	"fmt"
	"net/http"
	"projext/internal/driver"
	"strconv"
)

func SearchResult(w http.ResponseWriter, r *http.Request) {
	var result string
	id, _ := strconv.Atoi(r.FormValue("ExecutorId"))
	name := r.FormValue("Name")
	surname := r.FormValue("Surname")
	fatherName := r.FormValue("FatherName")
	position := r.FormValue("Position")

	println(name, surname, fatherName)
	output := driver.SearchByFio(id, name, surname, fatherName, position)

	for i, _ := range output {
		result += output[i]
		result += "\n\n"
	}
	//println(output)

	fmt.Fprintf(w, result)
}
