package order_activity

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"go.uber.org/cadence/activity"
)

func init() {
	activity.Register(NotifyCustomer)
}

func NotifyCustomer(ctx context.Context, orderID string, customerID string, content string) (string, error) {
	// activity.GetLogger(ctx).Info("SimpleActivity called.", zap.String("Value", value))
	spew.Dump("NotifyCustomer: order ", orderID, "customer", customerID, "content: ", content)
	return orderID, nil
}
