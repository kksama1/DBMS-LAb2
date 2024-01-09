package driver

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"projext/internal/model"
)

func SearchByFio(id int, name, surname, fatherName, position string) []string {
	findOptions := options.Find()
	filter := bson.D{{"executor", bson.D{
		{"id", id},
		{"firstName", name},
		{"lastName", surname},
		{"fatherName", fatherName},
		{"position", position},
	},
	},
	}
	client, ctx := CreateConnection()
	cur, _ := client.Database("ProjectManagment").Collection("Task").Find(context.TODO(), filter, findOptions)
	var results []string
	for cur.Next(context.TODO()) {
		var elem model.Task
		//Create a value into which the single document can be decoded
		fmt.Println("------")
		fmt.Println(cur.Current)

		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		out, err := json.Marshal(elem)
		if err != nil {
			panic(err)
		}

		results = append(results, string(out))

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	//Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents: %+v\n", results)
	defer closeConn(ctx, client)
	return results

}
