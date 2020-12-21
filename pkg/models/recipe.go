package models

import "time"

// Recipe struct
type Recipe struct {
	ID int
	Name string
	Ingredients []string
	Directions []string
	Category Category
	CookTime time.Time
	PrepTime time.Time
	Calories int
}