package main

import (
	"fmt"

	"entity"
	"html/template"
	"net/http"
)

//type producto struct {
//	nombre string
//	precio float64
//}

func test1(w http.ResponseWriter, r *http.Request) {
	p := entity.Producto{0, "Ungüento", 15.30}
	fmt.Fprint(w, p)

}

func showAdd(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("add_product.html"))
	t.Execute(w, nil)
}

func init() {
	http.HandleFunc("/", test1)
	http.HandleFunc("/add", showAdd)
}
