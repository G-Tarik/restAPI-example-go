package myapp

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller struct {
	Data Storage
}

func Authentication(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authToken := req.Header.Get("Authorization")
		if authToken != "" {
			rez := CheckToken(authToken)
			if rez {
				f(w, req)
			} else {
				json.NewEncoder(w).Encode(Exception{Message: "Token validation failed"})
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
	}
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	responseData, _ := json.MarshalIndent(endpointIndex, "", "    ")
	writeResponse(w, responseData)

	return
}

func (c *Controller) GetCurrencies(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	name := args["name"]

	currency := c.Data.GetCurrencies(name)
	responseData, _ := json.MarshalIndent(currency, "", "    ")
	writeResponse(w, responseData)

	return
}

func writeResponse(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
