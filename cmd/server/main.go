package main

import (
	"database/sql"
	"fmt"
	"github.com/FelipeAragao/must-have/configs"
	"github.com/FelipeAragao/must-have/docs"
	_ "github.com/FelipeAragao/must-have/docs"
	"github.com/FelipeAragao/must-have/internal/infra/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/swaggo/http-swagger"
	"net/http"
)

// @contact.name   Felipe Aragão
// @contact.email  felpe.thiago10@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
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

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Must Have API"
	docs.SwaggerInfo.Description = "This is a sample server Must Have server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

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
