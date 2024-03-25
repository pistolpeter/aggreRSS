package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	util "github.com/pistolpeter/aggreRSS/internal"
	"github.com/pistolpeter/aggreRSS/internal/database"
)

func handleReadinessGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := struct {
			Status string `json:"status"`
		}{
			Status: "ok",
		}
		util.RespondWithJSON(w, 200, resp)
		return
	}
}

func handleErr() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("CONNECTION")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Post("/users", handleUsersCreate(apiCfg.DB))
	v1Router.Get("/users", middlewareAuth(handleUsersGet, apiCfg.DB))

	v1Router.Post("/feeds", middlewareAuth(apiCfg.handleFeedsCreate, apiCfg.DB))
	v1Router.Get("/feeds", apiCfg.handleFeedsGetAll())

	v1Router.Post("/feed_follows", middlewareAuth(apiCfg.handleFeedFollowsCreate, apiCfg.DB))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.handleFeedFollowDelete())
	v1Router.Get("/feed_follows", middlewareAuth(apiCfg.handleFeedFollowGetAll, apiCfg.DB))

	v1Router.Get("/healthz", handleReadinessGet())
	v1Router.Get("/err", handleErr())

	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
