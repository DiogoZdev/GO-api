package main

import (
	"net/http"
	"root/controllers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/arts", controllers.NewArtHandlers().Get)
	mux.HandleFunc("/artists", controllers.NewArtistsHandle().Get)

	http.ListenAndServe(":3060", mux)
}