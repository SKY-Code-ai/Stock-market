package main

import (
	"fmt"
	"sync"

	"github.com/GaneshDeshmane/cex-matching-engine/internal/orderbook"
)

type Command struct {
	Order *orderbook.Order
}

func main() {
	inbox := make(chan Command)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for cmd := range inbox {
			fmt.Printf("Processing order: %+v\n", cmd.Order)
		}
	}()

	inbox <- Command{Order: &orderbook.Order{
		ID:        "1",
		Side:      orderbook.Buy,
		Type:      orderbook.Limit,
		Price:     100,
		Quantity:  10,
		Remaining: 10,
	}}

	close(inbox)
	wg.Wait()
}