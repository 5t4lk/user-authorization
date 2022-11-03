package signin

import (
	"123/internal/database"
	"go.mongodb.org/mongo-driver/bson"
)

func ChangeData() error {
	var userChange, refreshUserField, username, password string
	Input("What type of data you want to change?\n[1] - First name\n[2] - Last name\n[3] - Username\n[4] - Password\n[5] - Date of birth\n[6] - Balance", &userChange)

	switch userChange {
	case "1":
		userChange = "Firstname"
	case "2":
		userChange = "Lastname"
	case "3":
		userChange = "Username"
	case "4":
		userChange = "Password"
	case "5":
		userChange = "DateOfBirth"
	case "6":
		userChange = "Balance"
	}

	Input("Please enter again your username: ", &username)
	Input("Please enter again your password: ", &password)
	Input("Write new info", &refreshUserField)

	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		return err
	}
	defer database.Close(client, ctx, cancel)

	filter := bson.D{
		{"Username", bson.D{{"$eq", username}}},
		{"Password", bson.D{{"$eq", password}}},
	}

	update := bson.D{
		{"$set", bson.D{
			{userChange, refreshUserField},
		}},
	}

	_, err = database.UpdateOne(client, ctx, "UsersInfo", "data", filter, update)
	if err != nil {
		return err
	}

	return nil
}
