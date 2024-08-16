package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coderajay94/microservice1/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	MongoDatabase interface {
		Close(ctx context.Context) error
		SaveAccountDetails(model.UserResponseDB) (model.SaveResponseDB, error)
		GetAccountDetails(model.UserRequestDB) (model.UserResponseDB, error)
	}

	mongoDatabase struct{
		client *mongo.Client
		accountCollection *mongo.Collection
	}
)


func NewClient(host, username, password, dbname, collectionName string) (MongoDatabase, error){

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//mongodb://admin:admin@localhost:27017/?authSource=employees
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s/?authSource=%s", username, password, host, dbname)
	fmt.Println("connectionString", connectionString)
	//"mongodb://localhost:27017"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		fmt.Println("error while connecting to MongoDB", err)
		return nil, err
	}
	fmt.Println("connection established with database...")
	return &mongoDatabase{
		client: client,
		accountCollection: client.Database(dbname).Collection(collectionName),
	}, nil
}

func (db *mongoDatabase) Close(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}

func (md mongoDatabase) SaveAccountDetails(resp model.UserResponseDB) (model.SaveResponseDB, error) {

	_, err := md.accountCollection.InsertOne(context.TODO(), resp)
	if err != nil {
		fmt.Println("Error inserting account details")
		//panic(err)
		return model.SaveResponseDB{Status: "Failed to insert record"}, err
	}
	return model.SaveResponseDB{Status: "Success"}, nil
}

func(md mongoDatabase) GetAccountDetails(req model.UserRequestDB) (model.UserResponseDB,error){

	filter := bson.D{{Key: "email", Value: req.Email}}
	var resp model.UserResponseDB
	err := md.accountCollection.FindOne(context.TODO(), filter).Decode(&resp)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist:", req.Email)
		return model.UserResponseDB{}, err
	} else if err != nil {
		log.Fatal(err)
	}
	
	return resp, nil
}
