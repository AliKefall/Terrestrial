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
	SQLDB  *sql.DB
}

// Server constructor artık Config de alıyor
func newServer(sqlDB *sql.DB, cfg *Config) *Server {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	queries := db.New(sqlDB)

	s := &Server{
		Router: router,
		DB:     queries,
		Config: cfg,
		SQLDB:  sqlDB,
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

		r.Get("/monthly", handlers.SumByMonthHandler(s.SQLDB))

		r.Get("/daily", handlers.SumByDayHandler(s.SQLDB))

		r.Get("/yearly", handlers.SumByYearHandler(s.SQLDB))
	})
}
