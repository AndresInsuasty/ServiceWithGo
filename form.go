package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		fmt.Println(r.Method)
		r.ParseForm()
		fmt.Println("Nombre", r.Form["nombre"])
		fmt.Println("Codigo 1", r.Form["codigo1"])
		fmt.Println("Codigo 2", r.Form["codigo2"])
		fmt.Println("Detalle", r.Form["detalle"])
		fmt.Println("Proveedor", r.Form["proveedor"])
		fmt.Println("Responsable", r.Form["responsable"])
		fmt.Println("Presupuesto", r.Form["presupuesto"])
		fmt.Println("Fecha limite", r.Form["fechalimite"])
		fmt.Println("Periodicidad", r.Form["periodicidad"])
	}
}

func main() {
	http.HandleFunc("/login", login)
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css/"))))
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("img/"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server Error", err)
	}
}
