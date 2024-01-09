package driver

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type obj struct {
	Id      primitive.ObjectID `bson:"_id"`
	Course  string             `bson:"course"`
	Teacher string             `bson:"teacher"`
}

func closeConn(ctx context.Context, client *mongo.Client) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func CreateConnection() (*mongo.Client, context.Context) {
	uri := "mongodb://localhost:27017"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return client, ctx
}

func AddProject() {

	findOptions := options.Find()
	client, _ := CreateConnection()
	cur, _ := client.Database("IU6").Collection("Courses").Find(context.TODO(), bson.D{}, findOptions)
	var results []obj
	for cur.Next(context.TODO()) {
		var elem obj
		//Create a value into which the single document can be decoded
		fmt.Println(cur.Current)

		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(elem)

		results = append(results, elem)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	//Close the cursor once finished
	cur.Close(context.TODO())

	//output, _ := client.Database("IU6").ListCollectionNames(context.TODO(),
	//	bson.D{})
	fmt.Printf("Found multiple documents: %+v\n", results)
	//defer closeConn(ctx, client)
	//
	//client.Database("IU6").Collection("IU6-11M").Find(ctx, bson.D{})
	//fmt.Println(output)
}
