package routes

import (
  "Be_waysbeam/handlers"
  "Be_waysbeam/pkg/middleware"
  "Be_waysbeam/pkg/mysql"
  "Be_waysbeam/repositories"

  "github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
  productRepository := repositories.RepositoryProduct(mysql.DB)
  h := handlers.HandlerProduct(productRepository)

  r.HandleFunc("/products",h.FindProducts).Methods("GET") 
  r.HandleFunc("/product/{id}",middleware.Auth( h.GetProduct)).Methods("GET")
  r.HandleFunc("/product/{id}", middleware.Auth(h.DeleteProduct)).Methods("DELETE")
r.HandleFunc("/product", middleware.Auth(middleware.UploadFile(h.CreateProduct))).Methods("POST") 
r.HandleFunc("/product/{id}", middleware.Auth(middleware.UploadFile(h.UpdateProduct))).Methods("PATCH")
}