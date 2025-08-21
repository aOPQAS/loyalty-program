package models

type CreateProgramServiceRequest struct {
	ProgramID string `json:"program_id"`
	ServiceID string `json:"service_id"`
}
