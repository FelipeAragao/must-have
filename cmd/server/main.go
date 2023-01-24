package main

import (
	"github.com/FelipeAragao/must-have/configs"
	"github.com/FelipeAragao/must-have/internal/infra/server"
	"net/http"
)

func main() {
	config, err := configs.LoadConfig("./cmd/server/.env")
	if err != nil {
		panic(err)
	}
	//
	s := server.NewServer(
		config.JWTSecret,
		config.JwtExperesIn)
	r := s.Start()

	http.ListenAndServe(":3000", r)
}
