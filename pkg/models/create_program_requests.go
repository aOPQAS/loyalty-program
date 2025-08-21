package models

type CreateProgramRequest struct {
	Type              string `json:"type"`
	Name              string `json:"name"`
	Image             string `json:"image"`
	FixedPrice        int    `json:"fixed_price"`
	TotalServicesCost int    `json:"total_services_cost"`
	DiscountPercent   int    `json:"discount_percent"`
	ValidUntil        string `json:"valid_until"`
	Terms             string `json:"terms"`
	Active            bool   `json:"active"`
}
