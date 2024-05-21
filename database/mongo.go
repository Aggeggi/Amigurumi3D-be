package database

import (
	"log"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDb() {
	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoURL := os.Getenv("MONGO_URL")
	// Setup mgm default config
	credential := options.Credential{
		Username: mongoUsername,
		Password: mongoPassword,
	}
	err := mgm.SetDefaultConfig(nil, "amigurumi3d", options.Client().ApplyURI(mongoURL).SetAuth(credential))
	if err != nil {
		log.Fatal(err)
	}
}
