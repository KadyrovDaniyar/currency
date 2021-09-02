package server

import (
	"log"
	"projects/currency/app/models"
)

type R_currency interface {
	Create(items *models.Rates)
	CheckExists(date string) (exists bool, err error)
	FindWithDate(date string) (items *models.Rates, err error)
}

func(s *Server) Create(rates *models.Rates) {

	formattedDate,err := ChangeDateFormat(rates.Date)
	if err != nil {
		log.Println(err.Error())
	}

	for _, item := range rates.Items {

		if err := s.db.Create(&models.R_currency{
			Title: item.Fullname,
			Code: item.Title,
			Value: item.Description,
			A_date: formattedDate,
		}).Error; err != nil {
			log.Println(err.Error())
		}
	}
}

func(s *Server) CheckExists(date string) (exists bool, err error) {

	formattedDate,err := ChangeDateFormat(date)
	if err != nil {
		return false,err
	}

	var rate *models.R_currency

	if err = s.db.Where("A_date = ?", formattedDate).First(&rate).Error; err != nil {
		if err.Error() == "record not found"{
			err = nil
		}
		return false, err
	}

	if rate != nil {
		return true,nil
	}
	return false,nil
}

func(s *Server) FindWithDate(date string) (rates []*models.R_currency, err error) {
	formattedDate, err := ChangeDateFormat(date)
	if err != nil {
		return nil, err
	}

	if err = s.db.Where("A_date = ?", formattedDate).Find(&rates).Error; err != nil {
		return nil, err
	}

	return
}

func(s *Server) FindWithDateAndCode(date string, code string) (rates []*models.R_currency, err error) {
	formattedDate, err := ChangeDateFormat(date)
	if err != nil {
		return nil, err
	}

	if err = s.db.Where("A_date = ? and Code = ?", formattedDate, code).Find(&rates).Error; err != nil {
		return nil, err
	}

	return
}