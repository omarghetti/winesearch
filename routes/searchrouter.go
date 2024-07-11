package routes

import (
	"github.com/gorilla/mux"
	"github.com/oghetti/winesearch/services"
)

type SearchRouter struct {
	searchService *services.SearchService
}

func NewSearchRouter(searchService *services.SearchService) *SearchRouter {
	return &SearchRouter{searchService}
}

func (sr *SearchRouter) RegisterRoutes(r *mux.Router) {
	s := r.PathPrefix("/v1").Subrouter()
	s.HandleFunc("/search", sr.searchService.HandleSearch).Methods("GET")
}
