package server

import (
	web "htmx-templ-sqlc/cmd/web/templates"
	"log"
	"net/http"
)

var peopleTabs = []string{
	"All",
	"Interns",
	"Project Managers",
	"Admins",
}
var activeTab = "All"

func HelloWebHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	component := web.HelloPost(firstName, lastName)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	navbar := web.Navbar("home")
	component := web.Home(navbar)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}
func MediaHandler(w http.ResponseWriter, r *http.Request) {
	navbar := web.Navbar("media")
	component := web.Media(navbar)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in MediaHandler: %e", err)
	}
}

func PeopleHandler(w http.ResponseWriter, r *http.Request) {
	navbar := web.Navbar("people")
	tabComp := web.PeopleTabs(peopleTabs, activeTab)
	component := web.People(navbar, tabComp)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in PeopleHandler: %e", err)
	}
}
func PeopleTabHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	tabName := r.FormValue("tab")
	activeTab = tabName
	tabComp := web.PeopleTabs(peopleTabs, activeTab)
	err = tabComp.Render(r.Context(), w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in PeopleHandler: %e", err)
	}

}
func SummerProgramHandler(w http.ResponseWriter, r *http.Request) {
	navbar := web.Navbar("summer_program")
	component := web.SummerProgram(navbar)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in SummerHandler: %e", err)
	}
}
