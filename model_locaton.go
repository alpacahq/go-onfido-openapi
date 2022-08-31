package onfido_openapi

// Location struct for Location
type Location struct {
// The IP address of the applicant
IPAddress *string `json:"ip_address,omitempty"`
// The 3 character ISO country code of the country where the applicant resides
CountryOfResidence *string `json:"country_of_residence,omitempty"`
}