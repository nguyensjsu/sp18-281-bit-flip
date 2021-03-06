package model

import (
	"github.com/gocql/gocql"
	"time"
)

// Database Models
type StoreDetails struct {
	Id   gocql.UUID `json:"id" cql:"id"`
	Name string     `json:"name" cql:"name"`
}

type ProductDetails struct {
	Id   gocql.UUID `json:"id" cql:"id"`
	Name string     `json:"item" cql:"name"`
	Size string     `json:"size" cql:"size"`
}

type InventoryDetails struct {
	StoreId  gocql.UUID `json:"-" cql:"store_id"`
	Id       gocql.UUID `json:"id" cql:"id"`
	Name     string     `json:"item" cql:"name"`
	Quantity string     `json:"qty" cql:"quantity"`
	Size     string     `json:"size" cql:"size"`
}

type AllocationDetails struct {
	StoreId  gocql.UUID          `json:"store_id" cql:"store_id"`
	UserId   gocql.UUID          `json:"-" cql:"user_id"`
	Id       gocql.UUID          `json:"id" cql:"id"`
	Status   string              `json:"status" cql:"status"`
	Expires  time.Time           `json:"expires" cql:"expires"`
	Products []*InventoryDetails `json:"products" cql:"expires"`
}

// API Models
type ListStoresResult struct {
	Stores        []StoreDetails `json:"stores"`
	NextPageToken *gocql.UUID    `json:"next_page_token"`
}

type ListInventoryResult struct {
	Products []*InventoryDetails `json:"products"`
}

type ListAllocationsResult struct {
	Allocations []*AllocationDetails `json:"allocations"`
}

type CreateAllocationRequest struct {
	Products []*InventoryDetails `json:"products"`
}