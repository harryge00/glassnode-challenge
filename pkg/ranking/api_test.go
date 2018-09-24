package ranking

import (
	"testing"
)

func TestGetRanking(t *testing.T) {
	ranking, err := GetCoinList()
	if err != nil {
		t.Error(err)
	}
	t.Log(len(ranking.Data), ranking)
}
