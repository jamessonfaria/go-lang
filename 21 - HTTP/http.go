package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá mundo"))
}

func usuario(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregar pagina de usuarios..."))
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pagina raiz"))
	})

	http.HandleFunc("/home", home)
	http.HandleFunc("/usuario", usuario)

	log.Fatal(http.ListenAndServe(":5000", nil))
}