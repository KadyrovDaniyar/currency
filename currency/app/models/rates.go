package models

import (
	"encoding/xml"
)

type Rates struct {
	XMLName xml.Name `xml:"rates"`
	Items []Item `xml:"item"`
	Date string `xml:"date"`
}

type Item struct {
	Fullname string `xml:"fullname"`
	Title string `xml:"title"`
	Description float64 `xml:"description"`
}
