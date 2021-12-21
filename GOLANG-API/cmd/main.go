package main

import (
	"github.com/gorilla/mux"
	//"github.com/jinzhu/gorm/dialects/mysql"
	"log"

	"github.com/gorilla/handlers"

	"net/http"

	"github.com/khadyCi/bloober/pkg/routes"
)

type Dest struct {
	Name string
}

func main() {
	r := mux.NewRouter()
	//gestionar Errores de CORS
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Origin"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	s := r.PathPrefix("/api").Subrouter()
	routes.RegisterUserStoreRoutes(s)
	routes.RegisterTaskStoreRoutes(s)
	routes.RegisterTypeTaskStoreRoutes(s)
	routes.SendMailsRoutes(s)
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe("localhost:5555", handlers.CORS(header, methods, origins)(s)))

}
