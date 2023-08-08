package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

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

	fmt.Println("Running server...")
	app := fiber.New()
	var ctx = context.Background()

	app.Post("/", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			panic(err)
		}

		payload, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}

		if err := redisClient.Publish(ctx, "send-user-data", payload).Err(); err != nil {
			panic(err)
		}

		return c.SendStatus(200)
	})

	app.Listen(":3000")
}
