package services

import (
	"net/http"

	elastic "github.com/elastic/go-elasticsearch/v8"
)

type SearchService struct {
	client *elastic.TypedClient
}

func NewSearchService(client *elastic.TypedClient) *SearchService {
	return &SearchService{client}
}

func (ss *SearchService) HandleSearch(w http.ResponseWriter, r *http.Request) {
	// Handle search
	w.Write([]byte("Search"))
}
