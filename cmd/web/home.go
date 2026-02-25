package web

import (
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	navbar := Navbar("home")
	component := Home(navbar)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}
