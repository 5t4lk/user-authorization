package userAuthorization

type Info struct {
	Username    string `json:"username" bson:"username"`
	Password    string `json:"password" bson:"password"`
	Firstname   string `json:"firstname" bson:"firstname"`
	Lastname    string `json:"lastname" bson:"lastname"`
	DateOfBirth string `json:"dateOfBirth" bson:"dateOfBirth"`
	Balance     string `json:"balance" bson:"balance"`
}
