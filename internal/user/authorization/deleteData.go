package signin

import (
	"123/internal/database"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteData() error {
	var username, password string
	Input("Please enter your username again: ", &username)
	Input("Please enter your password again: ", &password)

	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		return err
	}
	defer database.Close(client, ctx, cancel)

	filter := bson.D{
		{"Username", username},
		{"Password", password},
	}

	_, err = database.DeleteOne(client, ctx, "UsersInfo", "data", filter)
	if err != nil {
		return err
	}

	return nil
}
