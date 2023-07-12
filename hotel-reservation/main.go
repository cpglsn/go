package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cpglsn/hotel-reservation/api"
	"github.com/cpglsn/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dburi    = "mongodb://localhost:27017"
	dbname   = "hotel-reservation"
	userColl = "users"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)
	user := types.User{
		FirstName: "James",
		LastName:  "Watercooler",
	}
	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	var James types.User
	if err := coll.FindOne(ctx, bson.M{}).Decode(&James); err != nil {
		log.Fatal(err)
	}

	fmt.Println(James)

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(":5000")
}

