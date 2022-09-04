package main

import (
	"Be_waysbeam/database"
	"Be_waysbeam/pkg/mysql"
	"Be_waysbeam/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	// "github.com/joho/godotenv"
)

// import "fmt"

func main() {
	//init database
	mysql.DatabaseInit()

	
	//run migration
	database.RunMigration()
	//gorila mux route
	r := mux.NewRouter()
	
	//path file
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer((http.Dir("./uploads")))))
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})


	var port = "5000"
	fmt.Println("server running localhost:"+port)
	
	// Embed the setup allowed in 2 parameter on this below code ...
	http.ListenAndServe("localhost:"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
