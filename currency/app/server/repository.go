package server

import (
	"fmt"
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

		query := fmt.Sprintf("Insert into R_Currency(Title, Code, Value, A_date) values (N'%s','%s',%v,'%s')",
			item.Fullname, item.Title, item.Description, formattedDate)

		_, err := s.db.Exec(query)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func(s *Server) CheckExists(date string) (exists bool, err error) {

	formattedDate,err := ChangeDateFormat(date)
	if err != nil {
		return false,err
	}

	query := fmt.Sprintf("select top 1 * from R_Currency where A_date ='%s'",
		formattedDate)

	result, err := s.db.Exec(query)
	if err != nil {
		if err.Error() == "record not found"{
			err = nil
		}
		return false, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rows > 0 {
		return true,nil
	}
	return false,nil
}

func(s *Server) FindWithDate(date string) (rates []*models.R_currency, err error) {
	formattedDate, err := ChangeDateFormat(date)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("select Title, Code, Value, A_date from R_Currency where A_date = '%s'",
		formattedDate)

	result, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		rate := models.R_currency{}

		err = s.db.QueryRow(query).Scan(&rate.Title, &rate.Code, &rate.Value, &rate.A_date)
		if err != nil {
			return nil, err
		}

		rates = append(rates,&rate)
	}

	return
}

func(s *Server) FindWithDateAndCode(date string, code string) (rates []*models.R_currency, err error) {
	formattedDate, err := ChangeDateFormat(date)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("select Title, Code, Value, A_date from R_Currency where A_date = '%s' and Code = '%s'",
		formattedDate, code)

	result, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		rate := models.R_currency{}

		err = s.db.QueryRow(query).Scan(&rate.Title, &rate.Code, &rate.Value, &rate.A_date)
		if err != nil {
			return nil, err
		}

		rates = append(rates,&rate)
	}

	return
}
