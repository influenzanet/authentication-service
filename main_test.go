package main

import (
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

var (
	testInstanceID = strconv.FormatInt(time.Now().Unix(), 10)
)

func dropTestDB() {
	log.Println("Drop test database")
	ctx, cancel := getContext()
	defer cancel()

	// Drop Test Instance Auth DB
	err := dbClient.Database(conf.DB.DBNamePrefix + "global-infos").Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
	result := m.Run()
	dropTestDB()
	os.Exit(result)
}
