package main

import (
	"fmt"
	"github.com/ini8labs/lsdb"
)

func main() {
	dbClient, err := lsdb.NewClient()
	if err != nil {
		panic(err.Error())
	}

	newUserInfo := lsdb.UserInfo{
		Name:  "Anand",
		Phone: 7506639417,
		GovID: "ABCDEFG",
		EMail: "anand@ini8labs.tech",
	}

	if err := dbClient.OpenConnection(); err != nil {
		panic(err.Error())
	}

	defer dbClient.CloseConnection()

	if err := dbClient.AddNewUserInfo(newUserInfo); err != nil {
		panic(err.Error())
	}
	fmt.Println("User Added Successfully")
}
