package orderbook
type OrderBook struct {
	Symbol string
	Bids   []*Order 
	Asks   []*Order 
}

func NewOrderBook(symbol string) *OrderBook {
	return &OrderBook{Symbol: symbol}
}
func (ob *OrderBook) Match(incoming *Order) []Trade {
	var trades []Trade

	if incoming.Side == Buy {
		trades = ob.matchAgainst(incoming, &ob.Asks, func(restingPrice int64) bool {
			return incoming.Price >= restingPrice 
		})
	} else {
		trades = ob.matchAgainst(incoming, &ob.Bids, func(restingPrice int64) bool {
			return incoming.Price <= restingPrice
		})
	}

	if incoming.Remaining > 0 && incoming.Type == Limit {
		ob.insertResting(incoming)
	}

	return trades
}

func (ob *OrderBook) matchAgainst(incoming *Order, opposite *[]*Order, crosses func(int64) bool) []Trade {
	var trades []Trade

	for incoming.Remaining > 0 && len(*opposite) > 0 {
		best := (*opposite)[0]

		if incoming.Type == Limit && !crosses(best.Price) {
			break
		}

		fillQty := min(incoming.Remaining, best.Remaining)

		buyID, sellID := incoming.ID, best.ID
		if incoming.Side == Sell {
			buyID, sellID = best.ID, incoming.ID
		}

		trades = append(trades, Trade{
			Price:       best.Price,
			Quantity:    fillQty,
			BuyOrderID:  buyID,
			SellOrderID: sellID,
		})

		incoming.Remaining -= fillQty
		best.Remaining -= fillQty

		if best.Remaining == 0 {
			*opposite = (*opposite)[1:]
		}
	}

	return trades
}

func (ob *OrderBook) insertResting(o *Order) {
	if o.Side == Buy {
		i := 0
		for i < len(ob.Bids) && ob.Bids[i].Price >= o.Price {
			i++
		}
		ob.Bids = append(ob.Bids, nil)
		copy(ob.Bids[i+1:], ob.Bids[i:])
		ob.Bids[i] = o
	} else {
		i := 0
		for i < len(ob.Asks) && ob.Asks[i].Price <= o.Price {
			i++
		}
		ob.Asks = append(ob.Asks, nil)
		copy(ob.Asks[i+1:], ob.Asks[i:])
		ob.Asks[i] = o
	}
}