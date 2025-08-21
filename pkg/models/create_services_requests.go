package models

type CreateServicesRequest struct {
	Name     string `json:"name"`
	Tarif    int    `json:"tarif"`
	Duration int    `json:"duration"`
}
