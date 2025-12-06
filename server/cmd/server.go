package main

import (
	"database/sql"

	"github.com/AliKefall/DonemOdevi/internal/db"
	"github.com/AliKefall/DonemOdevi/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
	DB     *db.Queries
	Config *Config
}

func newServer(sqlDB *sql.DB) *Server {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	queries := db.New(sqlDB)

	s := &Server{
		Router: router,
		DB:     queries,
	}

	s.registerRoutes()
	return s
}

func (s *Server) registerRoutes() {
	s.Router.Post("/auth/register", handlers.Register(s.DB))
	s.Router.Post("/auth/login", handlers.Login(s.DB, s.Config.TokenSecret))

	s.Router.Route("/transactions", func(r chi.Router) {
		r.Use(handlers.AuthMiddleware(s.Config.TokenSecret))
		r.Post("/", handlers.CreateTransaction(s.DB))
		r.Get("/", handlers.ListTransactions(s.DB))
		r.Get("/monthly", handlers.SumMonthly(s.DB))
		r.Get("/daily", handlers.SumDaily(s.DB))
		r.Get("/yearly", handlers.SumYearly(s.DB))
	})
}
