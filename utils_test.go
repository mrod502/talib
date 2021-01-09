package talib

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestDiff(t *testing.T) {}

func TestPercentDiff(t *testing.T) {
	arr := make([]decimal.Decimal, 10)

	for i := 0; i < 10; i++ {
		arr[i] = decimal.New(int64(i+1), 0)
	}

	fmt.Println(diffPercent(arr))

}
