package main

import (
	"client/server/db"
	"client/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"
)

type Server struct {
	NATS *nats.Conn
}

func NewServer(nt *nats.Conn) *Server {
	return &Server{
		NATS: nt,
	}
}

type CreateOrderInput struct {
	TenantID   int `json:"tenantID"`
	CustomerID int `json:"customerID"`
}

func (s *Server) CreateOrder(c *fiber.Ctx) error {
	input := &CreateOrderInput{}
	if err := c.BodyParser(input); err != nil {
		return err
	}

	order, err := db.CreateOrder(&db.Order{
		TenantID:   input.TenantID,
		CustomerID: input.CustomerID,
	})
	if err != nil {
		return err
	}
	// Publish a message to nats
	if _, err = s.NATS.Request(util.NATS_NEW_ORDER, []byte(order.ID), 1*time.Minute); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"isSuccess": true,
		"order":     order,
	})
}

type ListOrderInput struct {
	TenantID int    `json:"tenantID"`
	Status   string `json:"customerID"`
}

func (s *Server) ListOrders(c *fiber.Ctx) error {
	query := &ListOrderInput{}
	if err := c.QueryParser(query); err != nil {
		return err
	}

	orders := db.ListOrder(query.TenantID, query.Status)
	return c.JSON(fiber.Map{
		"isSuccess": true,
		"orders":    orders,
	})
}

func (s *Server) StartPreparing(c *fiber.Ctx) error {
	orderID := c.Params("id")
	// Publish a message to nats
	if _, err := s.NATS.Request(util.NATS_START_PREPARING_ORDER, []byte(orderID), 1*time.Minute); err != nil {
		return err
	}

	order, err := db.StartPreparingOrder(orderID)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"isSuccess": true,
		"order":     order,
	})
}

func (s *Server) PreparingOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")

	// Publish a message to nats
	if _, err := s.NATS.Request(util.NATS_PREPARING_ORDER, []byte(orderID), 1*time.Minute); err != nil {
		return err
	}

	order, err := db.PreparingOrder(orderID)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"isSuccess": true,
		"order":     order,
	})
}
