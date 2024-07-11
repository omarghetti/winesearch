package main

import (
	"log/slog"
	"net/http"

	elastic "github.com/elastic/go-elasticsearch/v8"
)

type App struct {
	Server        *http.Server
	ElasticClient *elastic.Client
	Logger        *slog.Logger
}
