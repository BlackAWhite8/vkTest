package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BlackAWhite8/vkTest/actors"
	"github.com/BlackAWhite8/vkTest/authorization"
	"github.com/BlackAWhite8/vkTest/movies"
	"github.com/gorilla/mux"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome page")
}

func authorizedWithPriv(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !authorization.Flag {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		} else if !authorization.Priv {
			http.Error(w, "access denied", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func authorizedWithoutPriv(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !authorization.Flag {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	/*без  параметров get у функций изменяющих бд можно посетить страницы без прав администратора
	но запросы отправить нельзя,не знаю как исправить без добавления get методов */
	r.Handle("/actors/update", authorizedWithPriv(http.HandlerFunc(actors.UpdateActor))).Methods("PUT", "GET")
	r.Handle("/actors/delete", authorizedWithPriv(http.HandlerFunc(actors.DeleteActor))).Methods("DELETE", "GET")
	r.Handle("/actors/create", authorizedWithPriv(http.HandlerFunc(actors.CreateActor))).Methods("POST", "GET")

	r.Handle("/actors", authorizedWithoutPriv(http.HandlerFunc(actors.GetActors))).Methods("GET")

	r.Handle("/movies/create", authorizedWithPriv(http.HandlerFunc(movies.CreateMovie))).Methods("POST", "GET")
	r.Handle("/movies/update", authorizedWithPriv(http.HandlerFunc(movies.UpdateMovie))).Methods("PUT", "GET")
	r.Handle("/movies/delete", authorizedWithPriv(http.HandlerFunc(movies.DeleteMovie))).Methods("DELETE", "GET")

	r.Handle("/movies/search", authorizedWithoutPriv(http.HandlerFunc(movies.SearchMovie))).Methods("GET")
	r.Handle("/movies", authorizedWithoutPriv(http.HandlerFunc(movies.GetMovies))).Methods("GET")

	r.HandleFunc("/login", authorization.Login).Methods("GET")
	r.HandleFunc("/signup", authorization.SignUp).Methods("POST")

	r.HandleFunc("/", mainPage)

	log.Println("Start server localhost:8080")
	err := (http.ListenAndServe(":8080", r))
	if err != nil {
		log.Println("server is down due to error")
		log.Fatal(err)
	}
}
