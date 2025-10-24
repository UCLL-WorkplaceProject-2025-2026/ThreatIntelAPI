package model

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
