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

type DrinkShort struct {
	gorm.Model
	Name string
	Price int
}
