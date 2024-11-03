package routes

import (
	"github.com/gorilla/mux"
	"github.com/harshith53/bookstore/pkg/controllers"
)

var RegisterBookstoreRouter = func(rounter *mux.Router) {

	rounter.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	rounter.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	rounter.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	rounter.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	rounter.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}
