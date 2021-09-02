package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"projects/currency/app/models"
)

// CurrencySave godoc
// @Summary Save currency to particular date
// @Description save currency by date, if exists on particular date then respond message already exists
// @Tags currency/save
// @Produce  json
// @Param date path string true "dd.mm.yyyy"
// @Success 200 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /currency/save/{date} [get]
func(s *Server) CurrencySave(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	date := vars["date"]

	err := CheckDateFormat(date)
	if err != nil {
		s.respond(w, r, map[string]interface{}{
			"success": false,
			"message": "Некорректный формат даты",
		}, 422)
		return
	}

	exists,err := s.CheckExists(date)
	if err != nil {
		log.Println(err)
		s.respond(w, r, map[string]interface{}{
			"success": false,
			"message": "Ошибка при проверке даты в бд",
		}, 422)
		return
	}

	if exists {
		s.respond(w, r, map[string]interface{}{
			"success": false,
			"message": "Запись на данную дату уже есть в бд",
		}, 422)
		return
	}

	rates, err := GetCurrencyByDate(date, s.config.CurrencyURL)
	if err != nil {
		log.Println(err)
		s.respond(w, r, map[string]interface{}{
			"success": false,
			"message": "Не смогли получить данные из нац.,банка",
		}, 422)
		return
	}

	go s.Create(rates)

	s.respond(w, r, map[string]interface{}{
		"success": true,
	}, 200)
	return
}

// CurrencyGet godoc
// @Summary Get all currency rates to particular date without code on concrete currency, if it exists on db
// @Description get currency to particular date without code on concrete currency
// @Tags currency
// @Produce json
// @Param date path string true "dd.mm.yyyy"
// @Success 200 {object} []models.Response
// @Failure 422 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /currency/{date} [get]
func(s *Server) CurrencyGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	date := vars["date"]

	err := CheckDateFormat(date)
	if err != nil {
		s.respond(w, r, map[string]interface{}{
			"success": false,
			"message": "Некорректный формат даты",
		}, 422)
		return
	}

	rates, err := s.FindWithDate(date)
	if err != nil {
		log.Println(err)
		s.respond(w, r, map[string]interface{}{
			"success": false,
			"message": "Проблемы при выборке с бд",
		}, 422)
		return
	}

	if len(rates) == 0 {
		s.respond(w, r, map[string]interface{}{
			"success": false,
			"message": "Не найдено",
		}, 404)
		return
	}

	response := make([]*models.Response,0)

	for _,value := range rates {
		response = append(response, &models.Response{
			Fullname: value.Title,
			Title: value.Code,
			Description: value.Value,
			A_date: value.A_date,
		})
	}

	s.respond(w, r, response,200)
}

// CurrencyGetWithCode godoc
// @Summary Get particular currency rate in particular date, if it exists on db
// @Description get currency to particular date without code on concrete currency
// @Tags currency
// @Produce json
// @Param date path string true "dd.mm.yyyy"
// @Param code path string true "3 digit string"
// @Success 200 {object} models.Response
// @Failure 422 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /currency/{date}/{code} [get]
func(s *Server) CurrencyGetWithCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	date := vars["date"]
	code := vars["code"]

	err := CheckDateFormat(date)
	if err != nil {
		s.respond(w, r, map[string]interface{}{
			"success": false,
			"message": "Некорректный формат даты",
		}, 422)
		return
	}

	if len(code) > 3 {
		s.respond(w, r, map[string]interface{}{
			"success": false,
			"message": "Некорректный формат кода",
		}, 422)
		return
	}

	rates, err := s.FindWithDateAndCode(date,code)
	if err != nil {
		log.Println(err)
		s.respond(w, r, map[string]interface{}{
			"success": false,
			"message": "Проблемы при выборке с бд",
		}, 422)
		return
	}

	if len(rates) == 0 {
		s.respond(w, r, map[string]interface{}{
			"success": false,
			"message": "Не найдено",
		}, 404)
		return
	}

	response := make([]*models.Response,0)

	for _,value := range rates {
		response = append(response, &models.Response{
			Fullname: value.Title,
			Title: value.Code,
			Description: value.Value,
			A_date: value.A_date,
		})
	}

	s.respond(w, r, response,200)
}
