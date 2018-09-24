package price

import (
	"testing"
)

func TestGetPrice(t *testing.T) {
	priceMap, err := GetPriceMap("", "10")
	if err != nil {
		t.Error(err)
	}
	t.Log(priceMap.PriceMap)
	if len(priceMap.PriceMap) != 10 {
		t.Fatalf("Lenght of priceMap should be 10, but is %v", len(priceMap.PriceMap))
	}
	priceMap, err = GetPriceMap("101", "20")
	if err != nil {
		t.Error(err)
	}
	t.Log(priceMap)
	if len(priceMap.PriceMap) != 20 {
		t.Fatalf("Lenght of priceMap should be 20, but is %v", len(priceMap.PriceMap))
	}
}
