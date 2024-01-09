package handlers

import (
	"fmt"
	"net/http"
	"projext/internal/driver"
)

func AddNewTasks(w http.ResponseWriter, r *http.Request) {
	driver.InsertTask()
	fmt.Fprintf(w, "Новый записи были добавлены")
}
