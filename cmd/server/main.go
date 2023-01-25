package main

import (
	"database/sql"
	"fmt"
	"github.com/FelipeAragao/must-have/configs"
	_ "github.com/FelipeAragao/must-have/docs"
	"github.com/FelipeAragao/must-have/internal/infra/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/swaggo/http-swagger"
	"net/http"
)

// @title           Must Have API
// @version         1.0
// @description     Must Have API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Felipe Aragão
// @contact.email  felipe.thiago10@gmail.com

// @host      localhost:3000
// @BasePath  /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig("./cmd/server/.env")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
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

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"), //TODO: change to env
	))
	http.ListenAndServe(configs.ServerPort, r)
}
