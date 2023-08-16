package main

import (
	"os"

	"github.com/FardeeUseng/t-shirt-backend/configs"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}

	cfg := new(configs.Configs)

	// fiber configs
	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_POST")

	// database configs
	cfg.PostgreSQL.Host = os.Getenv("DB_HOST")
	cfg.PostgreSQL.Port = os.Getenv("DB_POST")
	cfg.PostgreSQL.Protocal = os.Getenv("DB_PROTOCAL")
	cfg.PostgreSQL.Username = os.Getenv("DB_USERNAME")
	cfg.PostgreSQL.Password = os.Getenv("DB_PASSWORD")
	cfg.PostgreSQL.Database = os.Getenv("DB_DATABASE")
	cfg.PostgreSQL.SSLMode = os.Getenv("DB_SSL_MODE")

}
