package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var ctx = context.Background()

func main() {
	fmt.Println("Reading config...")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../config")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Create redisClient
	var host = viper.GetString("server.host")
	var port = viper.GetString("server.port")
	var redisClient = redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})

	fmt.Println("Running client...")

	subscriber := redisClient.Subscribe(ctx, "send-user-data")

	user := User{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal([]byte(msg.Payload), &user); err != nil {
			panic(err)
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", user)
	}
}
