package main

import (
	"net/http"

	"github.com/smapl/lifestyle-up/internal/time_management/handler"
)

func main() {
	http.HandleFunc("/time-management", handler.Handler)

}
