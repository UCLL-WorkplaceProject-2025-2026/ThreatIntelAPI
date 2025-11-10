package controller

import (
	"net/http"
)

type OpenPhishController struct {
	csvPath string
}

func NewOpenPhishController(csvPath string) *OpenPhishController {
	return &OpenPhishController{
		csvPath: csvPath,
	}
}

func (c *OpenPhishController) GetAll(w http.ResponseWriter, r *http.Request) {
	EnableCORS(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Serve the raw CSV file
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=\"feed.csv\"")
	http.ServeFile(w, r, c.csvPath)
}
