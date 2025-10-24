package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"threatintelapi/service"
)

type OpenPhishController struct {
	service service.OpenPhishService
}

func NewOpenPhishController(service service.OpenPhishService) *OpenPhishController {
	return &OpenPhishController{
		service: service,
	}
}

func (c *OpenPhishController) GetAll(w http.ResponseWriter, r *http.Request) {
	EnableCORS(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	records, err := c.service.GetAllRecords()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving records: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(records)
}
