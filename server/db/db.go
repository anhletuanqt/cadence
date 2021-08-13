package db

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/brianvoe/gofakeit/v6"
)

var OrdersTable = make(map[string]Order)

func init() {
	readFile()
}

func GetOrderByID(id string) Order {
	byteDate, _ := ioutil.ReadFile("../data.json")
	orderTable := make(map[string]Order)
	_ = json.Unmarshal(byteDate, &orderTable)
	OrdersTable = orderTable

	return OrdersTable[id]
}

func SetOrderByID(order Order) {
	OrdersTable[order.ID] = order
	writeToFile()
}

func writeToFile() {
	file, _ := json.MarshalIndent(OrdersTable, "", " ")
	_ = ioutil.WriteFile("../data.json", file, 0644)
}

func readFile() {
	byteDate, _ := ioutil.ReadFile("../data.json")
	orderTable := make(map[string]Order)
	_ = json.Unmarshal(byteDate, &orderTable)
	OrdersTable = orderTable
}

func CreateOrder(order *Order) (*Order, error) {
	id := gofakeit.UUID()
	if _, ok := OrdersTable[id]; ok {
		return nil, errors.New("the order already exists")
	}
	order.ID = id
	order.Status = "pending"
	order.Activities = []string{"new order"}
	OrdersTable[id] = *order
	writeToFile()
	return order, nil
}

func ListOrder(tenantID int, status string) []Order {
	readFile()
	orders := []Order{}

	for _, order := range OrdersTable {
		if tenantID != 0 && status != "" {
			if order.TenantID == tenantID && order.Status == status {
				orders = append(orders, order)
			}
		} else if tenantID != 0 {
			if order.TenantID == tenantID {
				orders = append(orders, order)
			}
		} else if status != "" {
			if order.Status == status {
				orders = append(orders, order)
			}
		} else {
			orders = append(orders, order)
		}

	}

	return orders
}

func StartPreparingOrder(id string) (*Order, error) {
	readFile()
	order, ok := OrdersTable[id]
	if !ok {
		return nil, errors.New("the order is not found")
	}
	order.Status = "start preparing"
	OrdersTable[id] = order
	writeToFile()
	return &order, nil
}

func PreparingOrder(id string) (*Order, error) {
	readFile()
	order, ok := OrdersTable[id]
	if !ok {
		return nil, errors.New("the order is not found")
	}
	order.Status = "preparing order"
	OrdersTable[id] = order
	writeToFile()
	return &order, nil
}
