package talib

import (
	"github.com/mrod502/util"
	"github.com/shopspring/decimal"
)

//Indicator - do we buy, sell, or hold?
type indicator byte

//CandlePart - OHLC?
type CandlePart byte

//constants
const (
	BUY  indicator = indicator('B')
	SELL indicator = indicator('S')
	HOLD indicator = indicator(0)
)

// ohlc
const (
	Open   = CandlePart(0)
	High   = CandlePart(1)
	Low    = CandlePart(2)
	Close  = CandlePart(3)
	Volume = CandlePart(4)
)

var (
	one     = decimal.New(1, 0)
	hundred = decimal.New(1, 2)
)

func getCandlePart(c util.Candle, p CandlePart) decimal.Decimal {
	switch p {
	case Open:
		return c.Open
	case High:
		return c.High
	case Low:
		return c.Low
	case Close:
		return c.Close
	case Volume:
		return c.Volume
	default:
		return c.Close
	}

}

func candlePartSlice(c []util.Candle, p CandlePart) (out []decimal.Decimal) {
	out = make([]decimal.Decimal, len(c))
	switch p {
	case Open:
		for k := range c {
			out[k] = c[k].Open
		}
	case High:
		for k := range c {
			out[k] = c[k].High
		}
	case Low:
		for k := range c {
			out[k] = c[k].Low
		}
	case Close:
		for k := range c {
			out[k] = c[k].Close
		}
	case Volume:
		for k := range c {
			out[k] = c[k].Volume
		}
	default:
		for k := range c {
			out[k] = c[k].Close
		}

	}

	return out
}

func splitPosNeg(s []decimal.Decimal) (pos []decimal.Decimal, neg []decimal.Decimal) {
	pos = make([]decimal.Decimal, 0, len(s)/2)
	neg = make([]decimal.Decimal, 0, len(s)/2)

	for _, d := range s {
		if d.LessThan(decimal.Zero) {
			neg = append(neg, d.Abs())
		}
		if d.GreaterThan(decimal.Zero) {
			pos = append(pos, d)
		}
	}

	return
}

func diff(d []decimal.Decimal) (out []decimal.Decimal) {
	out = make([]decimal.Decimal, len(d)-1)
	for ix := range out {
		out[ix] = d[ix+1].Sub(d[ix])
	}

	return
}

func diffPercent(d []decimal.Decimal) (out []decimal.Decimal) {
	out = make([]decimal.Decimal, len(d)-1)
	for ix := range out {
		out[ix] = hundred.Mul(d[ix+1].Sub(d[ix]).Div(d[ix]))
	}

	return
}
