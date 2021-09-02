package models

import "time"

type Response struct {
   Fullname string `json:"fullname"`
   Title string `json:"title"`
   Description float64 `json:"description"`
   A_date time.Time `json:"date"`
}
