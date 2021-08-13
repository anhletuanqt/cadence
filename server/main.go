package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"
)

func main() {
	natsConn, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		log.Fatal(err)
	}
	server := NewServer(natsConn)
	app := fiber.New()
	app.Post("/orders", server.CreateOrder)
	app.Get("/orders", server.ListOrders)
	app.Post("/orders/:id/start_preparing", server.StartPreparing)
	app.Post("/orders/:id/preparing_order", server.PreparingOrder)
	app.Listen(":8081")
}
