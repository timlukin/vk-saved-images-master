package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func getMatchHandler(w http.ResponseWriter, r *http.Request) {
	/*
		get parameter contains vk user_id
		returns matched images report
	*/
	fmt.Fprintln(w, "Not realized yet")
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/compare_request_form.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func serverErrorDecorator(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				http.Error(w, "Server failed", http.StatusInternalServerError)
			}
		}()
		f(w, r)
	}

}

func main() {
	http.HandleFunc("/", serverErrorDecorator(mainHandler))
	http.HandleFunc("/compare", serverErrorDecorator(getMatchHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
