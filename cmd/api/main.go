package main

import (
	"encoding/json"
	"net/http"

	"github.com/EvincarSN/go-init/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	//com chi
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/oder", Order)
	http.ListenAndServe(":8888", r)

	//puro
	// http.HandleFunc("/order", Order)
	//http.ListenAndServe(":8888", nil)
}

func Order(w http.ResponseWriter, r *http.Request) {
	// puro
	// if r.Method != http.MethodGet {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	json.NewEncoder(w).Encode("Method not allowed")
	// 	return
	// }

	order, err := entity.NewOrder("123", 1000, 10)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	order.CalculateFinalPrice()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}
