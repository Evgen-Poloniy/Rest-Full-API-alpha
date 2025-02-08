package main

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type CurrencySet map[string]float64

func CreateSet() CurrencySet {
	return CurrencySet{
		"RUB": 1.0,
		"USD": 97.28,
		"EUR": 103.50,
		"CNY": 13.3083,
		"JPY": 0.67,
		"KRW": 0.073,
		"GBP": 119.45,
		"INR": 1.17,
		"KZT": 0.21,
		"KGS": 1.11,
		"UZS": 0.0077,
		"TJS": 8.53,
		"UAH": 2.64,
		"BYN": 38.57,
		"TRY": 3.45,
	}
}

var currencySet CurrencySet

type CBRResponse struct {
	Valute map[string]struct {
		Value float64 `json:"Value"`
	} `json:"Valute"`
}

func updateExchangeRates(w http.ResponseWriter, money *float64, currency *string) bool {
	if *currency != "" {
		if *currency != "RUB" {
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			}

			client := &http.Client{
				Transport: tr,
				Timeout:   10 * time.Second,
			}

			url := "https://www.cbr-xml-daily.ru/daily_json.js"
			resp, err := client.Get(url)
			if err != nil {
				log.Println("Ошибка HTTP-запроса:", err)
				getResponseError(w, http.StatusNotFound, "Не удалось найти курс валют")
				return false
			}
			defer resp.Body.Close()

			var data CBRResponse
			err = json.NewDecoder(resp.Body).Decode(&data)
			if err != nil {
				getResponseError(w, http.StatusInternalServerError, "Не удалось декодировать курс валют")
				return false
			}

			if rate, exist := data.Valute[*currency]; exist {
				currencySet[*currency] = rate.Value
				*money *= rate.Value
			} else {
				getResponseError(w, http.StatusInternalServerError, "Не удалось выполнить конвертацию валют")
				return false
			}
		}
	} else {
		return false
	}

	return true
}
