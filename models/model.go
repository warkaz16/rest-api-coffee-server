package models

import "gorm.io/gorm"

type Drink struct {
	gorm.Model
	Name             string `json:"name"`
	Price            int    `json:"price"`
	InStock          bool   `json:"in_stock"`
	ContainsCaffeine bool   `json:"contains_caffeine"`
	Volume           int    `json:"volume"`
	Description      string `json:"description"`
}

type DrinkRequestDto struct {
	Name             *string `json:"name,omitempty"`
	Price            *int    `json:"price,omitempty"`
	InStock          *bool   `json:"in_stock,omitempty"`
	ContainsCaffeine *bool   `json:"contains_caffeine,omitempty"`
	Volume           *int    `json:"volume,omitempty"`
	Description      *string `json:"description,omitempty"`
}

type DrinkShort struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
