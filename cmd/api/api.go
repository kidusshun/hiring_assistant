package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	evaluationcritera "github.com/kidusshun/hiring_assistant/service/evaluation_critera"
	jobposting "github.com/kidusshun/hiring_assistant/service/job_posting"
	"github.com/kidusshun/hiring_assistant/service/resumes"
	"github.com/kidusshun/hiring_assistant/service/user"
)

type APIServer struct {
	addr string
	db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := chi.NewRouter()
	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000"}, // Allow specific origin
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Content-Type", "Authorization"},
			ExposedHeaders:   []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           300, // Max cache duration in seconds
		}),
	)

	router.Use(middleware.Logger)

	userStore := user.NewStore(s.db)
	userService := user.NewService(userStore)
	userHandler := user.NewHandler(userService)
	userHandler.RegisterRoutes(router)

	jobPostingStore := jobposting.NewStore(s.db)
	jobPostingService := jobposting.NewService(userStore, jobPostingStore)
	jobPostingHandler := jobposting.NewHandler(jobPostingService)
	jobPostingHandler.RegisterRoutes(router)

	evaluationCriteriaStore := evaluationcritera.NewStore(s.db)
	evaluationCriteriaService := evaluationcritera.NewService(evaluationCriteriaStore, userStore, jobPostingStore)
	evaluationCriteriaHandler := evaluationcritera.NewHandler(evaluationCriteriaService)
	evaluationCriteriaHandler.RegisterRoutes(router)

	resumeStore := resumes.NewStore(s.db)
	resumeService := resumes.NewService(resumeStore, userStore, jobPostingStore)
	resumeHandler := resumes.NewHandler(resumeService)
	resumeHandler.RegisterRoutes(router)
	
	log.Println("Listening on ", s.addr)
	err := http.ListenAndServe(s.addr, router)

	return err
}