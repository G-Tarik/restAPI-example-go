package myapp

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct{}

var db *sql.DB

func (r Storage) GetCurrencies(name string) PriceHistory {
	stmt := "select * from price_history"
	if name != "" {
		stmt += " WHERE currency = " + "'" + name + "'"
	}

	rows, err := db.Query(stmt)
	if err != nil {
		log.Fatal(err)
	}

	result := PriceHistory{}
	defer rows.Close()
	for rows.Next() {
		var r CurrencyPrice
		err = rows.Scan(&r.ID, &r.Currency, &r.BaseCurrency, &r.Price, &r.Time)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, r)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func CheckToken(token string) bool {
	var result int
	err := db.QueryRow("SELECT 1 FROM api_tokens WHERE token = $1", token).Scan(&result)
	if err != nil {
		log.Println(err)
	}
	if result == 1 {
		return true
	}
	return false
}

func ConnectDB(dbURI string) {
	var err error
	db, err = sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func DisconnectDB() {
	db.Close()
}
