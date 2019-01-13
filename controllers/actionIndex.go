package controllers

import (
	"crud/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func IndexPage(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		bks := models.GetBooks()
		tmpl, _ := template.ParseFiles("view/index.html")
		tmpl.Execute(w, bks)
	})
}

func Create(){
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case "GET":
			tmpl, _ := template.ParseFiles("view/create.html")
			tmpl.Execute(w, nil)
		case "POST":
			name := r.FormValue("name")
			err := models.SetBook(name)
			if err != nil {
				log.Println(err)
			}
			http.Redirect(w, r, "/", 301)
		default:
			fmt.Fprintf(w, "Метод не доступен")
		}
	})
}

func Update(){
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case "GET":
			id := r.FormValue("id")
			bk := models.GetBook(id)
			tmpl, _ := template.ParseFiles("view/update.html")
			tmpl.Execute(w, bk)
		case "POST":
			name := r.FormValue("name")
			id := r.FormValue("id")
			models.UpdateBook(id, name)
			http.Redirect(w, r, "/", 301)
		default:
			fmt.Fprintf(w, "Метод не доступен")
		}
	})
}

func Delete(){
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case "GET":
			id := r.FormValue("id")
			models.Del(id)
			http.Redirect(w, r, "/", 301)
		default:
			fmt.Fprintf(w, "Метод не доступен")
		}
	})
}




