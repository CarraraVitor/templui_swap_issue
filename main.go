package main

import (
	"context"
	"net/http"
	"tuissue/example"
)

func main() {
	app := http.NewServeMux()
	app.Handle(
		"GET /assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))),
	)

	app.HandleFunc("/", handleHome)
	app.HandleFunc("/value/{val}", handleValue)

	http.ListenAndServe("localhost:8000", app)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	example.PageExample().Render(context.Background(), w)
}

func handleValue(w http.ResponseWriter, r *http.Request) {
	val := r.PathValue("val")
	w.Write([]byte(val))
}
