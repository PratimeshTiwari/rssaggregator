package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/pratimeshtiwari/rssaggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Admission Portal !! ")
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbString := os.Getenv("DB_URL")
	if dbString == "" {
		log.Fatal("DB_URL variable is not set")
	}

	conn, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
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
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/error", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/courses", apiCfg.middlewareAuth(apiCfg.handlerCreateCourse))
	v1Router.Get("/courses", apiCfg.handlerGetCourses)
	v1Router.Post("/courses-enrolled", apiCfg.middlewareAuth(apiCfg.handlerCreateCourseEnrolled))
	v1Router.Get("/courses-enrolled", apiCfg.middlewareAuth(apiCfg.handlerGetCoursesEnrolled))
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Starting server on port %s", portString)
	srv.ListenAndServe()

	srverr := srv.ListenAndServe()
	if srverr != nil {
		log.Fatal("Error starting server:", srverr)
	}

	fmt.Println("PORT is set to:", portString)
}
