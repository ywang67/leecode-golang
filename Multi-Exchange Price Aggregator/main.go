package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Quote struct {
	Symbol string
	Price  float64
}

type Aggregator struct {
	mu   sync.RWMutex
	data map[string]float64
}

func NewAggregator(input <-chan Quote, done <-chan struct{}, wg *sync.WaitGroup) *Aggregator {
	agg := &Aggregator{data: make(map[string]float64)}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case q := <-input:
				agg.mu.Lock()
				agg.data[q.Symbol] = q.Price
				agg.mu.Unlock()
			case <-done:
				return
			}
		}
	}()

	return agg
}

func (a *Aggregator) Get(symbol string) (float64, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	p, ok := a.data[symbol]
	return p, ok
}

func exchange(name string, symbols []string, out chan<- Quote, stop <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stop:
				return
			default:
				symbol := symbols[rand.Intn(len(symbols))]
				price := 90 + rand.Float64()*20
				out <- Quote{Symbol: symbol, Price: price}
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
}

func main() {
	input := make(chan Quote, 100)
	done := make(chan struct{})
	stop := make(chan struct{})
	var wg sync.WaitGroup

	agg := NewAggregator(input, done, &wg)

	exchange("binance", []string{"BTC", "ETH"}, input, stop, &wg)
	exchange("coinbase", []string{"BTC", "SOL"}, input, stop, &wg)
	exchange("okx", []string{"ETH", "SOL"}, input, stop, &wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("\nSnapshot %d:\n", i+1)
			for _, sym := range []string{"BTC", "ETH", "SOL"} {
				if p, ok := agg.Get(sym); ok {
					fmt.Printf("%s: %.2f\n", sym, p)
				} else {
					fmt.Printf("%s: (no data)\n", sym)
				}
			}
		}
		close(stop)
		close(done)
	}()

	wg.Wait()
	fmt.Println("✅ Program exited cleanly.")
}
