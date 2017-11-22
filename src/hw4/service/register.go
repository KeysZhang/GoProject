package service

import (
	"net/http"
	"html/template"
	"hw4/models"
	"github.com/unrolled/render"
)

var errorMessageSlice []models.ErrorMessage

func register(w http.ResponseWriter, r *http.Request){
	r.ParseForm() 
	if r.Method == "POST"{
		user := &models.User{
			r.Form["name"][0],
			r.Form["id"][0],
			r.Form["phone"][0],
			r.Form["email"][0],
		}
		errorMessageSlice = models.IsAllValid(*user)
		if(len(errorMessageSlice) == 0){
			t, _ := template.ParseFiles("./view/Detail.html")
			t.Execute(w, *user)
		}else{
			t, _ := template.ParseFiles("./view/SignUp.html")
			t.Execute(w, *user)
		}
	}else{
		errorMessageSlice = nil
		t, _ := template.ParseFiles("./view/SignUp.html")
		t.Execute(w, nil)
	}
}

func errorCheck(formatter *render.Render) http.HandlerFunc{
	iserror := true
	if len(errorMessageSlice) == 0{
		iserror = true
	}
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			ErrorSlice      []models.ErrorMessage`json:"errorslice"`
			IsError         bool`json:"iserror"`
        }{ErrorSlice: errorMessageSlice, IsError: iserror})
    }
}
