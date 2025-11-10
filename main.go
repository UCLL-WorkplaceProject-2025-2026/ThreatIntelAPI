package main

import (
	"log"
	"net/http"
	"os"
	"threatintelapi/controller"
	"threatintelapi/repository"
	"threatintelapi/service"
)

func main() {
	csvPath := "resources/openphish/feed.csv"
	netcraftRepo := repository.NewNetcraftRepository("resources/netcraft/netcraft_mock_belgian.json")

	netcraftService := service.NewNetcraftService(netcraftRepo)

	openPhishController := controller.NewOpenPhishController(csvPath)
	netcraftController := controller.NewNetcraftController(netcraftService)
	healthController := controller.NewHealthController()

	http.HandleFunc("/api/v1/openphish", openPhishController.GetAll)
	http.HandleFunc("/api/v1/netcraft", netcraftController.GetAll)
	http.HandleFunc("/health", healthController.Health)

	// Serve OpenAPI spec
	http.HandleFunc("/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		http.ServeFile(w, r, "openapi.yaml")
	})

	// Serve Swagger UI docs
	http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		// redirect to /docs/ so relative assets resolve
		http.Redirect(w, r, "/docs/", http.StatusFound)
	})
	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("docs"))))

	port := "8000"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	// Start server
	log.Printf("Starting ThreatIntel API server on port %s", port)
	log.Printf("Endpoints:")
	log.Printf("  - GET /api/v1/openphish")
	log.Printf("  - GET /api/v1/netcraft")
	log.Printf("  - GET /health")
	log.Printf("  - GET /openapi.yaml")
	log.Printf("  - GET /docs (Swagger UI)")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
