package signin

import (
	"123/internal/database"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func SignIn() (string, string, error) {
	var username, password string
	Input("Enter your username: ", &username)
	Input("Enter your password: ", &password)

	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		return "", "", err
	}
	defer database.Close(client, ctx, cancel)

	var filter, option interface{}

	filter = bson.D{
		{"Username", username},
		{"Password", password},
	}

	cursor, err := database.Query(client, ctx, "UsersInfo", "data", filter, option)

	var results []bson.D
	if err = cursor.All(ctx, &results); err != nil {
		return "", "", err
	}
	if results == nil {
		fmt.Print("User is not found. Make sure the input is correct.\n")
		SlowDown()
		SignIn()
	} else {
		fmt.Print("Successfully logged in!\n")
		return username, password, nil
	}

	return "", "", nil
}
