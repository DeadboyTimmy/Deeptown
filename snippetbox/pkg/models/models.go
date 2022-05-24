package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Product struct {
	ID          int
	Name        string
	Price       float64
	Vendor      string
	Category    string
	Subcategory string
	Sold        int
	Left        int
	G_from      string
	G_way       string
	G_to        string
	Trust_Level int
}
