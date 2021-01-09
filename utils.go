package talib

//BuySellHold - do we buy, sell, or hold?
type BuySellHold byte

//constants
const (
	BUY  = byte('B')
	SELL = byte('S')
	HOLD = byte('H')
)
