package conf

import (
	"jrpc-orm/models"
	"jrpc-orm/modules"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	modules.InitAllModules()
	models.ConnectDatabase(os.Getenv("DATABASE_DSN"))
}
