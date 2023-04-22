package main

import (
	"fmt"
	"github.com/ini8labs/lsdb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func main() {
	dbClient, err := lsdb.NewClient()
	if err != nil {
		panic(err.Error())
	}

	if err := dbClient.OpenConnection(); err != nil {
		panic(err.Error())
	}

	defer dbClient.CloseConnection()

	newEvent := lsdb.LotteryEventInfo{
		Name:          "Monday Special",
		EventDate:     primitive.NewDateTimeFromTime(time.Date(2023, time.April, 25, 0, 0, 0, 0, time.Local)),
		EventType:     "MS",
		WinningNumber: 1,
	}

	if err := dbClient.AddNewEvent(newEvent); err != nil {
		panic(err.Error())
	}

	fmt.Println("Added event successfully")
}
