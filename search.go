package main

import (
	"encoding/json"
	"net/http"
)

func (a *App) HandleSearch(w http.ResponseWriter, r *http.Request) {
	a.Logger.Info("Handling search request")
	winePayload := NewWine(
		"Château de Sales",
		"Pomerol",
		"Red",
		"Good",
		90,
		45.0,
	)
	enco := json.NewEncoder(w)
	enco.Encode(winePayload)
}
