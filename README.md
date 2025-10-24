# ThreatIntel API

A RESTful API service for threat intelligence data from OpenPhish and Netcraft sources.

## Architecture

This project follows a clean architecture pattern with clear separation of concerns:

```
ThreatIntelAPI/
├── main.go                 # Application entry point
├── model/                  # Data models/entities
│   ├── openphish.go       # OpenPhish record structure
│   └── netcraft.go        # Netcraft record structure
├── repository/            # Data access layer
│   ├── openphish_repository.go
│   └── netcraft_repository.go
├── service/               # Business logic layer
│   ├── openphish_service.go
│   └── netcraft_service.go
├── controller/            # HTTP handlers/routes
│   ├── openphish_controller.go
│   ├── netcraft_controller.go
│   ├── health_controller.go
│   └── utils.go
└── resources/             # Data files
    ├── openphish/
    │   └── feed.csv
    └── netcraft/
        └── netcraft_mock_belgian.json
```

## Layers

### 1. **Model Layer** (`model/`)
- Defines data structures and entities
- Contains `OpenPhishRecord` and `NetcraftRecord` structs
- Pure data models with no business logic

### 2. **Repository Layer** (`repository/`)
- Handles data access and file I/O operations
- Abstracts data sources (CSV, JSON files)
- Implements repository interfaces
- Responsibilities:
  - Read CSV/JSON files
  - Parse and convert raw data to models
  - Error handling for file operations

### 3. **Service Layer** (`service/`)
- Contains business logic
- Orchestrates data flow between repository and controller
- Implements service interfaces
- Responsibilities:
  - Call repository methods
  - Apply business rules (if any)
  - Logging
  - Data validation

### 4. **Controller Layer** (`controller/`)
- Handles HTTP requests and responses
- Maps HTTP requests to service calls
- Responsibilities:
  - Request validation
  - CORS handling
  - HTTP response formatting
  - Error response handling

## API Endpoints

### Get OpenPhish Data
```
GET /api/v1/openphish
```
Returns all OpenPhish threat intelligence records in JSON format.

```go
type OpenPhishRecord struct {
	URL             string `json:"url"`
	Brand           string `json:"brand"`
	IP              string `json:"ip"`
	ASN             string `json:"asn"`
	ASNName         string `json:"asn_name"`
	CountryCode     string `json:"country_code"`
	CountryName     string `json:"country_name"`
	TLD             string `json:"tld"`
	DiscoverTime    string `json:"discover_time"`
	FamilyID        string `json:"family_id"`
	Host            string `json:"host"`
	ISOTime         string `json:"isotime"`
	PageLanguage    string `json:"page_language"`
	SSLCertIssuedBy string `json:"ssl_cert_issued_by"`
	SSLCertIssuedTo string `json:"ssl_cert_issued_to"`
	SSLCertSerial   string `json:"ssl_cert_serial"`
	IsSpear         string `json:"is_spear"`
	Sector          string `json:"sector"`
}

```

### Get Netcraft Data
```
GET /api/v1/netcraft
```
Returns all Netcraft Belgian threat intelligence records in JSON format.

```go
type NetcraftRecord struct {
	ID                   string      `json:"id"`
	GroupID              string      `json:"group_id"`
	AttackURL            string      `json:"attack_url"`
	ReportedURL          string      `json:"reported_url"`
	IP                   string      `json:"ip"`
	CountryCode          string      `json:"country_code"`
	DateSubmitted        string      `json:"date_submitted"`
	LastUpdated          string      `json:"last_updated"`
	Region               string      `json:"region"`
	TargetBrand          string      `json:"target_brand"`
	AuthGiven            bool        `json:"authgiven"`
	Host                 string      `json:"host"`
	Registrar            string      `json:"registrar"`
	CustomerLabel        string      `json:"customer_label"`
	DateAuthed           interface{} `json:"date_authed"`
	StopMonitoringDate   string      `json:"stop_monitoring_date"`
	Domain               string      `json:"domain"`
	Language             string      `json:"language"`
	DateFirstActioned    string      `json:"date_first_actioned"`
	Escalated            bool        `json:"escalated"`
	FirstContact         string      `json:"first_contact"`
	FirstInactive        string      `json:"first_inactive"`
	IsRedirect           string      `json:"is_redirect"`
	AttackType           string      `json:"attack_type"`
	DeceptiveDomainScore float64     `json:"deceptive_domain_score"`
	DomainRiskRating     float64     `json:"domain_risk_rating"`
	FinalResolved        string      `json:"final_resolved"`
	FirstResolved        string      `json:"first_resolved"`
	Hostname             string      `json:"hostname"`
	EvidenceURL          string      `json:"evidence_url"`
	DomainAttack         string      `json:"domain_attack"`
	FalsePositive        bool        `json:"false_positive"`
	HostnameAttack       string      `json:"hostname_attack"`
	ReportSource         string      `json:"report_source"`
	Reporter             string      `json:"reporter"`
	ScreenshotURL        string      `json:"screenshot_url"`
	Status               string      `json:"status"`
}
```

### Health Check
```
GET /health
```
Returns API health status.

## OpenAPI documentation

- Spec file: `openapi.yaml`
- Served at runtime: `GET /openapi.yaml`

You can view it in your browser with an online editor:

1. Open https://editor.swagger.io/
2. Click File > Import URL, then paste: `http://localhost:8000/openapi.yaml` (adjust the port if you changed `PORT`).

Or use Redocly (optional):

```bash
npx redoc-cli serve openapi.yaml
```

### In-app Swagger UI

You can also view interactive docs served by the app itself:

- Swagger UI: `GET /docs` (served from `docs/index.html` and points to `/openapi.yaml`)

When running in Docker, the same paths are available inside the container since `openapi.yaml` and `docs/` are copied into the image.

## Running the Application

1. Docker:

```bash
docker build --no-cache -t threatintelapi .
docker run -p 8000:8000 threatintelapi
```

2. Using go (if golang is installed):

```bash
# Run the application
go run main.go

# Or build and run
go build -o threatintelapi
./threatintelapi
```

The API will start on port `8000` by default. You can override this by setting the `PORT` environment variable:

```bash
PORT=3000 go run main.go
```

## Dependencies

This project uses only Go standard library packages:
- `encoding/csv` - CSV parsing
- `encoding/json` - JSON encoding/decoding
- `net/http` - HTTP server
- `io` - I/O operations
- `os` - OS operations
- `log` - Logging

## Testing the API

```bash
# Test OpenPhish endpoint
curl http://localhost:8080/api/v1/openphish

# Test Netcraft endpoint
curl http://localhost:8080/api/v1/netcraft

# Test health endpoint
curl http://localhost:8080/health
```

## Architecture Benefits

1. **Separation of Concerns**: Each layer has a single, well-defined responsibility
2. **Testability**: Each layer can be tested independently with mocks
3. **Maintainability**: Changes to one layer don't affect others
4. **Scalability**: Easy to add new data sources or endpoints
5. **Dependency Injection**: Dependencies are injected, making the code flexible
