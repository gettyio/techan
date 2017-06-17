package talib4g

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewTradingRecord(t *testing.T) {
	record := NewTradingRecord()

	assert.Len(t, record.Trades, 0)
	assert.True(t, record.CurrentTrade().IsNew())
}

func TestTradingRecord_CurrentTrade(t *testing.T) {
	record := NewTradingRecord()

	yesterday := time.Now().Add(-time.Hour * 24)

	record.Enter(NM(1, USD), NS(2), yesterday)

	assert.EqualValues(t, 1, record.CurrentTrade().EntranceOrder().Price.Float())
	assert.EqualValues(t, 2, record.CurrentTrade().EntranceOrder().Amount.Float())
	assert.EqualValues(t, yesterday.UnixNano(),
		record.CurrentTrade().EntranceOrder().ExecutionTime.UnixNano())

	now := time.Now()
	record.Exit(NM(3, USD), NS(4), now)
	assert.True(t, record.CurrentTrade().IsNew())

	lastTrade := record.LastTrade()

	assert.EqualValues(t, 3, lastTrade.ExitOrder().Price.Float())
	assert.EqualValues(t, 4, lastTrade.ExitOrder().Amount.Float())
	assert.EqualValues(t, now.UnixNano(),
		lastTrade.ExitOrder().ExecutionTime.UnixNano())
}