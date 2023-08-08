package routes

import {
	"github.com/gorilla/mux"
	"github.com/kanayofirms/bookstore-management-api/pkg/controllers"
}

var RegisterBookStore = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
}