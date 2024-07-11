package main

import (
	"log"
	"net/http"
	"os"

	elastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/oghetti/winesearch/routes"
	"github.com/oghetti/winesearch/services"
)

type App struct {
	Server   *http.Server
	EsClient *elastic.TypedClient
}

func buildNewServer() *http.Server {
	router := mux.NewRouter().StrictSlash(true)

	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
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
	client, err := elastic.NewTypedClient(elastic.Config{
		Addresses: []string{
			getEnv("ELASTICSEARCH_URL"),
		},
	})

	if err != nil {
		log.Fatalf("Error creating the ES client: %s", err)
	}

	server := buildNewServer()

	app := &App{
		Server:   server,
		EsClient: client,
	}

	// Register routes
	searchService := services.NewSearchService(app.EsClient)
	searchRouter := routes.NewSearchRouter(searchService)
	searchRouter.RegisterRoutes(app.Server.Handler.(*mux.Router))

	log.Fatal(app.Server.ListenAndServe())

}
