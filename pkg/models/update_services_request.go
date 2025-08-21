package models

type UpdateServicesRequest struct {
	ServiceID string `json:"service_id"`
	Name      string `json:"name"`
	Tarif     int    `json:"tarif"`
	Duration  int    `json:"duration"`
}
