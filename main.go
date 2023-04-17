package main

import (
	"Merchant-Bank/delivery"

	_ "github.com/lib/pq"
)

func main() {
	delivery.Server().Run()
}
