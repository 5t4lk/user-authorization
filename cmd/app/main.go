package main

import (
	"123/internal/database"
	signin "123/internal/user/authorization"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func main() {
	username, password, err := signin.AskUserStart()
	if err != nil {
		log.Fatal(err)
	}

	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close(client, ctx, cancel)

	var filter interface{}

	filter = bson.D{
		{"Username", username},
		{"Password", password},
	}

	cursor, err := database.Query(client, ctx, "UsersInfo", "data", filter, nil)
	if err != nil {
		log.Fatal(err)
	}

	var results []bson.D

	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	if results != nil {
		fmt.Println(results)
	}

	signin.SlowDown()

	err = signin.AskUserChange()
	if err != nil {
		log.Fatal(err)
	}

	if results != nil {
		fmt.Println(results)
	}
}
