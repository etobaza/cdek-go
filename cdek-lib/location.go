package cdek_lib

type Location struct {
	Code        string `json:"code,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	City        string `json:"city,omitempty"`
	Address     string `json:"address,omitempty"`
}
