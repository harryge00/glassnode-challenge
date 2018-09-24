package ranking

import (
	"testing"
)

func TestGetPrice(t *testing.T) {
	ranking, err := GetPriceMap("", "10")
	if err != nil {
		t.Error(err)
	}
	t.Log(ranking)
	if len(ranking) != 10 {
		t.Fatalf("Lenght of rankingData should be 10, but is %v", len(ranking))
	}
	ranking, err = GetPriceMap("101", "20")
	if err != nil {
		t.Error(err)
	}
	t.Log(ranking)
	if len(ranking) != 20 {
		t.Fatalf("Lenght of rankingData should be 20, but is %v", len(ranking))
	}
}
