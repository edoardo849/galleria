package storage

import (
	"context"
	"log"
	"os"

	gStorage "cloud.google.com/go/storage"
)

var (
	// DB is the image database
	DB ImageDatabase

	// StorageBucket holds the handle for the bucket
	StorageBucket *gStorage.BucketHandle

	// StorageBucketName is the name of the bucket
	StorageBucketName string
)

func init() {
	var err error

	DB = newMemoryDB()

	// DB, err = configureCloudSQL(cloudSQLConfig{
	// 	Username: "admin",
	// 	Password: "admin",
	// 	// The connection name of the Cloud SQL v2 instance, i.e.,
	// 	// "project:region:instance-id"
	// 	// Cloud SQL v1 instances are not supported.
	// 	Instance: "",
	// })

	// TODO create bucket
	StorageBucketName = "bcg-progimage"
	StorageBucket, err = configureStorage(StorageBucketName)
	if err != nil {
		log.Fatal(err)
	}
}

func configureStorage(bucketID string) (*gStorage.BucketHandle, error) {
	ctx := context.Background()
	client, err := gStorage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.Bucket(bucketID), nil
}

type cloudSQLConfig struct {
	Username, Password, Instance string
}

func configureCloudSQL(config cloudSQLConfig) (ImageDatabase, error) {

	if os.Getenv("GAE_INSTANCE") != "" {
		// Running in production.
		return newMySQLDB(MySQLConfig{
			Username:   config.Username,
			Password:   config.Password,
			UnixSocket: "/cloudsql/" + config.Instance,
		})
	}

	// Running locally.
	return newMySQLDB(MySQLConfig{
		Username: config.Username,
		Password: config.Password,
		Host:     "localhost",
		Port:     3306,
	})
}
