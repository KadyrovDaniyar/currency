package server

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"projects/currency/app/models"
	"time"
)

const (
	layout = "02.01.2006"
	layoutDb = "2006-01-02"
)

func GetCurrencyByDate(date, urlTemplate string) (rates *models.Rates, err error) {
	url := fmt.Sprintf(urlTemplate,date)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	defer resp.Body.Close()

	byteValue, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}

	err = xml.Unmarshal(byteValue, &rates)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}

	return
}

func CheckDateFormat(date string) (err error) {
	_, err = time.Parse(layout,date)
	if err != nil {
		return
	}
	return
}

func ChangeDateFormat(date string) (changeDate string, err error) {
	t, err := time.Parse(layout,date)
	changeDate = t.Format(layoutDb)

	return
}
