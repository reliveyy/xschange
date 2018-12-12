package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/emj365/xschange/models"
	"github.com/gorilla/mux"
)

var orders = []models.Order{}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/orders", postOrders).Methods("POST")
	log.Println("server is running on 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func extractOrderFromRequest(r *http.Request, o *models.Order) {
	json.NewDecoder(r.Body).Decode(o)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(orders)
}

func postOrders(w http.ResponseWriter, r *http.Request) {
	var o models.Order
	extractOrderFromRequest(r, &o)
	o.Create(&orders)
	json.NewEncoder(w).Encode(o)
}
