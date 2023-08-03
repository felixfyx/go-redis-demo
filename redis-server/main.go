package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func main() {
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
