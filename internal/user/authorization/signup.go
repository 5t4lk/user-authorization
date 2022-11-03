package signin

import (
	userAuthorization "123"
	"123/internal/database"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func Signup() error {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		return err
	}
	defer database.Close(client, ctx, cancel)

	var userInfo userAuthorization.Info

	Input("Enter your first name: ", &userInfo.Firstname)
	Input("Enter your last name: ", &userInfo.Lastname)
	Input("Enter your date of birth: ", &userInfo.DateOfBirth)
	Input("Enter your username: ", &userInfo.Username)
	Input("Enter your password: ", &userInfo.Password)
	Input("Enter your balance: ", &userInfo.Balance)

	var doc interface{}

	doc = bson.D{
		{"Username", userInfo.Username},
		{"Password", userInfo.Password},
		{"Firstname", userInfo.Firstname},
		{"Lastname", userInfo.Lastname},
		{"DateOfBirth", userInfo.DateOfBirth},
		{"Balance", userInfo.Balance},
	}

	_, err = database.InsertOne(client, ctx, "UsersInfo", "data", doc)
	if err != nil {
		return err
	}

	fmt.Print("Successfully signed up.\n")
	return nil
}
