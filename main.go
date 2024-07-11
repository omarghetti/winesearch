package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	elastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func (a *App) RegisterRoutes() {
	r := mux.NewRouter().StrictSlash(true)
	v1api := r.PathPrefix("/api/v1").Subrouter()
	v1api.HandleFunc("/search", a.HandleSearch).Methods("GET")
	a.Server.Handler = r
}

func buildNewServer() *http.Server {

	s := &http.Server{
		Addr: ":8080",
	}

	return s
}

func getEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)

}

func main() {
	client, err := elastic.NewClient(elastic.Config{
		Addresses: []string{
			getEnv("ELASTICSEARCH_URL"),
		},
	})

	if err != nil {
		log.Fatalf("Error creating the ES client: %s", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &App{
		Server:        buildNewServer(),
		ElasticClient: client,
		Logger:        logger,
	}

	app.RegisterRoutes()

	log.Fatal(app.Server.ListenAndServe())

}
