package models

import "time"

type R_currency struct {
	Id int `gorm:"primaryKey"`
	Title string
	Code string
	Value float64
	A_date time.Time
}

func (c R_currency) TableName() string {
	return "R_CURRENCY"
}
