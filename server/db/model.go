package db

type Order struct {
	ID         string
	CustomerID int
	TenantID   int
	Status     string
	Activities []string
}
