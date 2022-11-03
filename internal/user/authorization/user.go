package signin

import (
	"fmt"
	"os"
	"time"
)

func AskUserStart() (string, string, error) {
	var userChoice string
	Input("To continue using our application, you need to be authorized.\n[1] - Sign-in\n[2] - Sign-up", &userChoice)

	if userChoice == "1" {
		username, password, err := SignIn()
		if err != nil {
			return "", "", err
		}
		return username, password, err
	} else if userChoice == "2" {
		err := Signup()
		if err != nil {
			return "", "", err
		}
	} else {
		fmt.Print("Incorrect input.\n")
		SlowDown()
		AskUserStart()
	}

	return "", "", nil
}

func AskUserChange() error {
	var userChoice string
	Input("Do you want to change something in your data?\n[1] - Change Data\n[2] - Delete Data\n[3] - Exit", &userChoice)
	if userChoice == "1" {
		if err := ChangeData(); err != nil {
			return err
		}
		fmt.Print("Successfully updated.\n")
		return nil
	} else if userChoice == "2" {
		if err := DeleteData(); err != nil {
			return err
		}
		fmt.Print("Successfully deleted.\n")
		return nil
	} else if userChoice == "3" {
		os.Exit(0)
	} else {
		fmt.Print("Incorrect input.\n")
		SlowDown()
		AskUserChange()
	}

	return nil
}

func Input(text, value interface{}) {
	fmt.Println(text)
	fmt.Scan(value)
}

func SlowDown() {
	for i := 0; i < 5; i++ {
		time.Sleep(444 * time.Millisecond)
		fmt.Print("* ")
		if i == 4 {
			fmt.Print("\n")
		}
	}
}
