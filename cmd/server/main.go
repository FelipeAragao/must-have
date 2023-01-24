package main

import (
	"database/sql"
	"fmt"
	"github.com/FelipeAragao/must-have/configs"
	"github.com/FelipeAragao/must-have/internal/infra/server"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	configs, err := configs.LoadConfig("./cmd/server/.env")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	s := server.NewServer(
		db,
		configs.JWTSecret,
		configs.JwtExperesIn,
	)
	r := s.Start()

	http.ListenAndServe(configs.WebServerPort, r)
}
