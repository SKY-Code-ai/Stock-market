package orderbook

import "testing"

func TestMatchExactFill(t *testing.T) {
	ob := NewOrderBook("BTC-USD")

	sell := &Order{ID: "sell-1", Side: Sell, Type: Limit, Price: 100, Quantity: 10, Remaining: 10}
	ob.Match(sell) // no opposite orders yet, rests on the book

	buy := &Order{ID: "buy-1", Side: Buy, Type: Limit, Price: 100, Quantity: 10, Remaining: 10}
	trades := ob.Match(buy)

	if len(trades) != 1 {
		t.Fatalf("expected 1 trade, got %d", len(trades))
	}
	if trades[0].Quantity != 10 {
		t.Errorf("expected quantity 10, got %d", trades[0].Quantity)
	}
	if trades[0].Price != 100 {
		t.Errorf("expected price 100, got %d", trades[0].Price)
	}
	if len(ob.Asks) != 0 {
		t.Errorf("expected asks to be empty after full fill, got %d", len(ob.Asks))
	}
}

func TestMatchPartialFill(t *testing.T) {
	ob := NewOrderBook("BTC-USD")

	sell := &Order{ID: "sell-1", Side: Sell, Type: Limit, Price: 100, Quantity: 10, Remaining: 10}
	ob.Match(sell)

	buy := &Order{ID: "buy-1", Side: Buy, Type: Limit, Price: 100, Quantity: 4, Remaining: 4}
	trades := ob.Match(buy)

	if len(trades) != 1 || trades[0].Quantity != 4 {
		t.Fatalf("expected 1 trade of qty 4, got %+v", trades)
	}
	if len(ob.Asks) != 1 || ob.Asks[0].Remaining != 6 {
		t.Fatalf("expected resting sell to have 6 remaining, got %+v", ob.Asks)
	}
}

func TestNoMatchWhenPricesDontCross(t *testing.T) {
	ob := NewOrderBook("BTC-USD")

	sell := &Order{ID: "sell-1", Side: Sell, Type: Limit, Price: 100, Quantity: 10, Remaining: 10}
	ob.Match(sell)

	buy := &Order{ID: "buy-1", Side: Buy, Type: Limit, Price: 95, Quantity: 10, Remaining: 10}
	trades := ob.Match(buy)

	if len(trades) != 0 {
		t.Fatalf("expected no trades, got %d", len(trades))
	}
	if len(ob.Bids) != 1 {
		t.Fatalf("expected buy order to rest on the book, got %d bids", len(ob.Bids))
	}
}