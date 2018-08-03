package myapp

import (
	"time"
)

type Exception struct {
	Message string `json:"message"`
}

type Tokens struct {
	Token string `json:"token"`
}

type CurrencyPrice struct {
	ID           int       `bson:"_id"`
	Currency     string    `json:"currency"`
	BaseCurrency string    `json:"base_currency"`
	Price        float64   `json:"price"`
	Time         time.Time `json:"time"`
}

type PriceHistory []CurrencyPrice
