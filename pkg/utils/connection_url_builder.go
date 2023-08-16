package utils

import (
	// "backend-t-shirt/configs"
	"errors"
	"fmt"

	"github.com/FardeeUseng/t-shirt-backend/configs"
)

func ConnectionUrlBuilder(stuff string, cfg *configs.Configs) (string, error) {
	var url string

	switch stuff {
	case "fiber":
		url = fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port)
	case "postgressql":
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.PostgresSQL.Host,
			cfg.PostgresSQL.Port,
			cfg.PostgresSQL.Username,
			cfg.PostgresSQL.Password,
			cfg.PostgresSQL.Database,
			cfg.PostgresSQL.SSLMode,
		)
	default:
		errMsg := fmt.Sprintf("error, connection url builder doesn't know the %s", stuff)
		return "", errors.New(errMsg)
	}
	return url, nil
}
