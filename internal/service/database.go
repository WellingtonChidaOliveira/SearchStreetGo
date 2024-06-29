package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/wellingtonchida/searchstreet/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	ConsultingCep
	ConsultingCepFromDb
	InsertCep
}

type ConsultingCepFromDb interface {
	GetCepFromDb(context.Context, string) (types.Address, error)
}
type InsertCep interface {
	InsertCep(context.Context, *types.Address) error
}

type CepFetcher struct {
	db *mongo.Client
}

var (
	host           = os.Getenv("DB_HOST")
	port           = os.Getenv("DB_PORT")
	databaseName   = os.Getenv("DB_DATABASENAME")
	collectionName = os.Getenv("DB_COLLECTION")
	username       = os.Getenv("DB_USERNAME")
	password       = os.Getenv("DB_PASS")
)

func Init() Service {
	log.Printf("DB_HOST: %s, DB_PORT: %s, DB_DATABASENAME: %s, DB_COLLECTION: %s, DB_USERNAME: %s, DB_PASS: %s",
		host, port, databaseName, collectionName, username, password)
	uri := fmt.Sprintf("mongodb://" + host + ":" + port)
	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.Auth = &options.Credential{
		Username: username,
		Password: password,
	}

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return &CepFetcher{
		db: client,
	}
}

func (c *CepFetcher) GetCepFromDb(ctx context.Context, cep string) (types.Address, error) {
	filter := bson.D{{"zipcode", cep}}
	addressmodel := types.Address{}
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	collection := c.db.Database(databaseName).Collection(collectionName)

	err := collection.FindOne(ctx, filter).Decode(&addressmodel)
	if err == mongo.ErrNoDocuments {
		log.Println("CEP not found")
		return types.Address{}, nil
	} else if err != nil {
		return addressmodel, err
	}

	return addressmodel, nil
}

func (c *CepFetcher) InsertCep(ctx context.Context, address *types.Address) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	collection := c.db.Database(databaseName).Collection(collectionName)

	_, err := collection.InsertOne(ctx, address)
	if err != nil {
		return err
	}

	return nil
}
