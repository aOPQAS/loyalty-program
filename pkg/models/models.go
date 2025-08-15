package models

type Program struct {
	ID                string  `json:"id" db:"id"` // UUID
	Type              string  `json:"type" db:"type"`
	Name              string  `json:"name" db:"name"`
	Image             string  `json:"image" db:"image"`
	FixedPrice        int     `json:"fixed_price" db:"fixed_price"`
	TotalServicesCost int     `json:"total_services_cost" db:"total_services_cost"`
	DiscountPercent   int     `json:"discount_percent" db:"discount_percent"`
	ValidUntil        string  `json:"valid_until" db:"valid_until"`
	Terms             string  `json:"terms" db:"terms"`
	CreatedAt         float64 `json:"created_at" db:"created_at"`
	UpdatedAt         float64 `json:"updated_at" db:"updated_at"`
	Active            bool    `json:"active" db:"active"`
}

type Service struct {
	ServiceID string `json:"service_id" db:"service_id"` // UUID
	Name      string `json:"name" db:"name"`
	Tarif     int    `json:"tarif" db:"tarif"`
	Duration  int    `json:"duration" db:"duration"`
}

type ProgramService struct {
	ProgramID string `json:"program_id" db:"program_id"` // UUID
	ServiceID string `json:"service_id" db:"service_id"` // UUID
}
