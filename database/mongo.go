package database

import (
	"context"
	"fmt"
	"time"

	retry "github.com/sethvargo/go-retry"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InputConnected struct {
	url string
}

func NewConnected(url string) InputConnected {
	return InputConnected{url: url}
}

type MongodbInterface interface {
	InitConnected() (*mongo.Client, error)
}

func (input InputConnected) InitConnected() (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI(input.url)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {

		ctx := context.Background()

		b := retry.NewFibonacci(5 * time.Second)
		var i int = 0
		for i < 2 {

			if err := retry.Do(ctx, retry.WithMaxRetries(2, b), func(ctx context.Context) error {
				if i == 1 {
					clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
					_, err := mongo.Connect(context.TODO(), clientOptions)
					fmt.Println("IN LOOP I==1")
					if err != nil {

						fmt.Println(" i ==1 error", i)
						return err
					}
					fmt.Println("Success")
					return nil
				}
				_, err := mongo.Connect(context.TODO(), clientOptions)
				fmt.Println("retry")
				if err != nil {

					fmt.Println(" err != nil", i)
					i++
					return err
				}

				fmt.Println("---")
				return nil
			}); err == nil {
				// handle error
				clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
				client, _ := mongo.Connect(context.TODO(), clientOptions)
				fmt.Println(" Err != nil ", "s")
				return client, nil
			}
		}

	}

	// Check the connection
	err = client.Ping(context.Background(), nil)

	if err != nil {

		return nil, err
	}

	log.Info("MongoClient connected")

	return client, nil
}

/*
func (db Mongodb) Connected(url string) (*mongo.Client, error) {
	fmt.Println(db.url)
	clientOptions := options.Client().ApplyURI(db.url)

	// Connect to MongoDB

	client, err := db.mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)

	if err != nil {
		return nil, err
	}

	log.Info("MongoClient connected")

	return client, nil
}

func (mongodb Mongodb) reconnect(url string) {
	err := retry.Do(func() error {
		client, err := mongo.Connect(context.TODO(), mongodb.url, mongodb.mongo.ConnectTimeout(5*time.Second))
		if err != nil {
			log.Printf("Failed to connect to MongoDB at %s: %s", url, err)
			return err
		}
		log.Println("Reconnected successfully.")
		mongodb.mongo = client
		return nil
	},
		retry.Attempts(3),
		retry.Delay(4*time.Second),
		retry.OnRetry(func(n uint, err error) {
			log.Printf("Retry %d: %s", n, err)
		}),
	)
	if err != nil {
		log.Fatal("Problem in connecting MongoDB.. exiting..")
	}
}
*/
