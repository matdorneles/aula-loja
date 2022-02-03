package main

import (
	"net/http"

	"github.com/matheus/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
