package orderbook

type Side int

const (
	Buy Side = iota
	Sell
)

type OrderType int

const (
	Limit OrderType = iota
	Market
)

type Order struct {
	ID        string
	Side      Side
	Type      OrderType
	Price     int64
	Quantity  int64
	Remaining int64
}

type Trade struct {
	Price      int64
	Quantity   int64
	BuyOrderID string
	SellOrderID string
}