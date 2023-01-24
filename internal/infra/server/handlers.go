package server

import (
	"github.com/FelipeAragao/must-have/internal/usecase/user/user_verifier"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/oauth"
	"net/http"
	"time"
)

type Server struct {
	Router       chi.Router
	JWTSecret    string
	JWTExperesIn int
}

func NewServer(secret string, experesIn int) *Server {
	return &Server{
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

	oauthServer := startOauth(s.JWTSecret, s.JWTExperesIn)
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	s.Router.Post("/authenticate/sso", oauthServer.UserCredentials)
	s.Router.Post("/authenticate", oauthServer.ClientCredentials)

	return s.Router
}

func startOauth(secret string, jwtExpiresIn int) *oauth.BearerServer {
	return oauth.NewBearerServer(
		secret,
		time.Second*time.Duration(jwtExpiresIn),
		user_verifier.NewUserVerifier(nil),
		nil)
}
