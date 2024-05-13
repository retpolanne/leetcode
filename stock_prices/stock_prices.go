package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

type Transaction struct {
	Profit int
	BoughtDate int
	PriceAtDay int
}

func (t *Transaction) Init() {
	t.Profit = 0
	t.PriceAtDay = -1
	t.BoughtDate = -1
}

func (t *Transaction) Buy (day int, prices []int) error {
	if day >= len(prices) {
		return fmt.Errorf("day %d out of bounds - len %d\n", day, len(prices))
	}
	t.BoughtDate = day
	t.PriceAtDay = prices[day]
	return nil
}

func (t *Transaction) Sell (day int, prices []int) error {
	if day < t.BoughtDate {
		return fmt.Errorf("can't sell stock at a past date - bought at %d, selling at %d!", t.BoughtDate, day)
	}
	if t.BoughtDate < 0 {
		return fmt.Errorf("you need to buy the stock before selling it! Transaction %v\n", &t)
	}
	t.Profit = prices[day] - t.PriceAtDay
	return nil
}

func maxProfit(prices []int) int {
	return -1
}
