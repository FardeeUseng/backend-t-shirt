package databases

import (
	"log"

	"github.com/FardeeUseng/t-shirt-backend/configs"
	"github.com/FardeeUseng/t-shirt-backend/pkg/utils"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPostgresSQLDBConnection(cfg *configs.Configs) (*sqlx.DB, error) {
	postgresUrl, err := utils.ConnectionUrlBuilder("postgressql", cfg)

	if err != nil {

		return nil, err
	}

	db, err := sqlx.Connect("pgx", postgresUrl)
	if err != nil {
		defer db.Close()
		log.Printf("error, can't connect to database, %s", err.Error())
		return nil, err
	}

	log.Println("postgresSQL database has been connected")
	return db, nil
}
