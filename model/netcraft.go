package model

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
