package todolists

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func createTodolist(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "todolist created\n path: %s", r.URL.Path)
}

func updateTodolist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "todolist with id: %s updated", id)
}

func InitTodolistsRouter(parentRouter *mux.Router) {
	todolistsRouter := parentRouter.PathPrefix("/todolists").Subrouter()

	todolistsRouter.HandleFunc("/", createTodolist).Methods("POST")
	todolistsRouter.HandleFunc("/{id}", updateTodolist).Methods("PATCH")

	parentRouter.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		log.Printf("Route: %s, Methods: %v", path, methods)
		return nil
	})
}
