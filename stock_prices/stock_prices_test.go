package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, 5, maxProfit([]int{7,1,5,3,6,4}))
}

func TestNoProfit(t *testing.T) {
	assert.Equal(t, 0, maxProfit([]int{7,6,4,3,1}))
}

func TestBuySellStock(t *testing.T) {
	transaction := &Transaction{}
	transaction.Init()
	prices := []int{7,1,5,3,6,4}
	assert.NoError(t, transaction.Buy(1, prices))
	assert.NoError(t, transaction.Sell(4, prices))
	assert.Equal(t, 5, transaction.Profit)
	assert.Equal(t, 1, transaction.PriceAtDay)
}

func TestBuySellStockAtLoss(t *testing.T) {
	transaction := &Transaction{}
	transaction.Init()
	prices := []int{7,1,5,3,6,4}
	assert.NoError(t, transaction.Buy(0, prices))
	assert.NoError(t, transaction.Sell(3, prices))
	assert.Equal(t, -4, transaction.Profit)
	assert.Equal(t, 7, transaction.PriceAtDay)
}

func TestBuySellStockPastDate(t *testing.T) {
	transaction := &Transaction{}
	transaction.Init()
	prices := []int{7,1,5,3,6,4}
	assert.NoError(t, transaction.Buy(1, prices))
	assert.Error(t, transaction.Sell(0, prices))
}

func TestSellStockBeforeBuying(t *testing.T) {
	transaction := &Transaction{}
	transaction.Init()
	prices := []int{7,1,5,3,6,4}
	assert.Error(t, transaction.Sell(0, prices))
}
