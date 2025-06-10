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

	app.HandleFunc("/", handlePage)
	app.HandleFunc("/swap", handleSwap)
	http.ListenAndServe("localhost:8000", app)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	example.PageExample().Render(context.Background(), w)
}

func handleSwap(w http.ResponseWriter, r *http.Request) {
	example.Section().Render(context.Background(), w)
}
