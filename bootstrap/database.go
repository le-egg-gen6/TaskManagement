package bootstrap

import (
	"context"
	"fmt"
	"go-ecommerce/mongo"
	"log"
	"time"
)

func NewMongoDatabase(env *Env) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPassword := env.DBPassword

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPassword, dbHost, dbPort)

	if dbUser == "" || dbPassword == "" {
		mongoURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}

	client, err := mongo.NewClient(mongoURI)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("MongoDB connection closed")
}
