package myapp

import (
	"net/http"
)

var controller = &Controller{Data: Storage{}}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
    {
        "Index",
        "GET",
        "/",
        controller.Index,
    },
    {
        "GetAllCurrencies",
        "GET",
        "/currency",
        Authentication(controller.GetCurrencies),
    },
    {
        "GetCurrency",
        "GET",
        "/currency/{name}",
        controller.GetCurrencies,
    },
}

type Endpoint struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Path   string `json:"path"`
}

var endpointIndex []Endpoint
