package order_activity

import (
	"client/server/db"
	"context"

	"go.uber.org/cadence/activity"
)

func init() {
	activity.Register(UpdateLCD)
}

func UpdateLCD(ctx context.Context, orderID string) (string, error) {
	// get order
	order := db.GetOrderByID(orderID)
	order.Activities = append(order.Activities, "update lcd")
	// set order
	db.SetOrderByID(order)
	return orderID, nil
}
