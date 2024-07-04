package api

import (
	"encoding/json"
	"github.com/Simnek/web-go/types"
	"github.com/Simnek/web-go/util"
	"net/http"
)

func HandleGetCustomer(w http.ResponseWriter, r *http.Request) {
	colName := r.URL.Query().Get("colName")
	value := r.URL.Query().Get("value")
	tableName := r.URL.Query().Get("tableName")

	result, err := util.SingleRowQuery(colName, value, tableName)
	if err != nil {
		// handle the error
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	switch result.(type) {
	case types.Customer:
		customer := result.(types.Customer)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	case types.Order:
		order := result.(types.Order)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(order)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "unknown result type"})
	}
}
