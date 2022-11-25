package util

import "github.com/profclems/go-dotenv"

func IsProduction() bool {
	return dotenv.GetString("RAILWAY_ENVIRONMENT") == "production"
}
