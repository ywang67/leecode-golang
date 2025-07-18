package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	data := []string{"aaa", "bbb", "ccc"}
	ch := make(chan string, len(data))
	wg := new(sync.WaitGroup)
	maxWorkers := 2

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for i := 0; i < maxWorkers; i++ {
		go worker(ctx, ch, wg)
	}

	for _, val := range data {
		wg.Add(1)
		ch <- val
	}

	close(ch)
	wg.Wait()
}

func worker(ctx context.Context, ch chan string, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("when ctx done because of ctx timeout, ctx.Done() it returns a channel and has nothing")
			return
		case record, ok := <-ch:
			if !ok {
				return
			}

			defer wg.Done()
			time.Sleep(2 * time.Second)
			fmt.Println("show channel element: ", record)
		}
	}
}
