package server_test

import (
	"testing"

	"microservice/internal/server"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePriceAndDiscount(t *testing.T) {
	tests := []struct {
		name            string
		fixedPrice      int
		totalCost       int
		discountPercent int
		wantPrice       int
		wantDiscount    int
	}{
		{
			name:            "fixedPrice > 0",
			fixedPrice:      80,
			totalCost:       100,
			discountPercent: 0,
			wantPrice:       80,
			wantDiscount:    20,
		},
		{
			name:            "discountPercent > 0, fixedPrice = 0",
			fixedPrice:      0,
			totalCost:       200,
			discountPercent: 25,
			wantPrice:       150,
			wantDiscount:    25,
		},
		{
			name:            "both zero",
			fixedPrice:      0,
			totalCost:       120,
			discountPercent: 0,
			wantPrice:       120,
			wantDiscount:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrice, gotDiscount := server.CalculatePriceAndDiscount(tt.fixedPrice, tt.totalCost, tt.discountPercent)
			assert.Equal(t, tt.wantPrice, gotPrice)
			assert.Equal(t, tt.wantDiscount, gotDiscount)
		})
	}
}
