package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"routing-gorilla-mux/todolists"
)

func main() {
	router := mux.NewRouter()
	todolists.InitTodolistsRouter(router)
	http.ListenAndServe(":8080", router)

}
