package handlers

import (
	"fmt"
	"net/http"
	"projext/internal/driver"
	"projext/internal/model"
	"strconv"
	"time"
)

func PostProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("Id"))
	fmt.Println(r.FormValue("Code"))
	code, _ := strconv.Atoi(r.FormValue("Code"))
	sp, _ := strconv.Atoi(r.FormValue("Sp"))
	fmt.Println(r.FormValue("Sp"))

	template := "2006-01-02T15:04"
	deadline, _ := time.Parse(template, r.FormValue("Deadline"))
	fmt.Println(r.FormValue("Deadline"))
	projectInfo := model.Project{
		Name:     r.FormValue("Id"),
		Code:     code,
		Sp:       sp,
		Deadline: deadline,
	}
	driver.CreateProject(projectInfo)

}
