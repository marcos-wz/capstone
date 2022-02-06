package entity

import "time"

type Fruit struct {
	ID           int       `json:"id" validate:"required"`            // Required, non-zero
	Name         string    `json:"name" validate:"required,alphanum"` // Required, only letters and numbers allowed
	Description  string    `json:"description"`                       // Optional
	Color        string    `json:"color" validate:"required,alpha"`   // Required, only letters allowed
	Unit         string    `json:"unit" validate:"oneof=kg lb"`       // Required, one of: kgs, lbs
	Price        float64   `json:"price"`                             // Optional, the price per unit
	Stock        int       `json:"stock"`                             // Optional, Items availability
	CaducateDays int       `json:"caducate_days"`                     // Optional, the number days to spend the fruit to be caducated
	Country      string    `json:"country" validate:"required,alpha"` // Rquired, the origin country only letters allowed
	CreatedAt    time.Time `json:"created_at" validate:"required"`    // Required
}
