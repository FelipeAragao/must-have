package server

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/oauth"
	"time"
)

type Server struct {
	Router       chi.Router
	JWTSecret    string
	JWTExperesIn int
	DB           *sql.DB
}

func NewServer(db *sql.DB, secret string, experesIn int) *Server {
	return &Server{
		DB:           db,
		Router:       chi.NewRouter(),
		JWTSecret:    secret,
		JWTExperesIn: experesIn,
	}
}

func (s *Server) Start() chi.Router {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "HEAD", "OPTION"},
		AllowedHeaders:   []string{"User-Agent", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "DNT", "Host", "Origin", "Pragma", "Referer"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	routerOauth(s)
	routerUser(s)

	return s.Router
}

func routerUser(s *Server) {
	s.Router.Route("/api/v1/users", func(r chi.Router) {
		r.Post("/", InitializeUserHandler(s.DB).CreateUser)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(oauth.Authorize(s.JWTSecret, nil))
			r.Put("/change-password", InitializeUserHandler(s.DB).ChangePassword)
			r.Put("/", InitializeUserHandler(s.DB).UpdateUser)
			r.Get("/", InitializeUserHandler(s.DB).GetUserById)
		})
	})
}

func routerOauth(s *Server) {
	oauthServer := oauth.NewBearerServer(
		s.JWTSecret,
		time.Second*time.Duration(s.JWTExperesIn),
		InitializeUserVerifier(s.DB),
		nil)

	s.Router.Route("/api/v1/authenticate", func(r chi.Router) {
		r.Post("/sso", oauthServer.ClientCredentials)
		r.Post("/", oauthServer.UserCredentials)
	})
}
