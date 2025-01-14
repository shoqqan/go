package todolists

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func createTodolist(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "todolist created\n path: %s", r.URL.Path)
	if err != nil {
		return
	}
}

func updateTodolist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	_, err := fmt.Fprintf(w, "todolist with id: %s updated", id)
	if err != nil {
		return
	}
}

func InitTodolistsRouter(parentRouter *mux.Router) {
	todolistsRouter := parentRouter.PathPrefix("/todolists").Subrouter()

	todolistsRouter.HandleFunc("", createTodolist).Methods("POST").Schemes("http")
	todolistsRouter.HandleFunc("{id}", updateTodolist).Methods("PATCH").Schemes("http")

	err := parentRouter.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		log.Printf("Route: %s, Methods: %v", path, methods)
		return nil
	})
	if err != nil {
		return
	}
}
