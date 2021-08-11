package order_activity

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"go.uber.org/cadence/activity"
)

func init() {
	activity.Register(UpdateLCD)
}

func UpdateLCD(ctx context.Context, orderID string, status string) (string, error) {
	// activity.GetLogger(ctx).Info("SimpleActivity called.", zap.String("Value", value))
	spew.Dump("UpdateLCD: order ", orderID, "with status", status)
	return orderID, nil
}
