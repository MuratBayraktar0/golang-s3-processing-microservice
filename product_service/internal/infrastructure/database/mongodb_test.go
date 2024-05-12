package database_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"product_service/internal/infrastructure/database"
)

func TestNewMongoDB(t *testing.T) {
	// Set up test variables
	uri := "mongodb://localhost:27017"
	dbName := "test_db"

	Convey("Given a MongoDB URI and database name", t, func() {
		mongoDB, err := database.NewMongoDB(uri, dbName)

		Convey("It should create a new MongoDB instance without error", func() {
			So(err, ShouldBeNil)
			So(mongoDB, ShouldNotBeNil)
			So(mongoDB.DB, ShouldNotBeNil)
		})

		Convey("It should disconnect from the MongoDB server without error", func() {
			err = mongoDB.Disconnect()
			So(err, ShouldBeNil)
		})
	})
}
