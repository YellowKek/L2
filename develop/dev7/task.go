package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var or func(channels ...<-chan any) <-chan any
	or = orChannel
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Millisecond),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(1*time.Millisecond),
		sig(1*time.Millisecond),
	)

	fmt.Printf("fone after %v", time.Since(start))

}

func orChannel(channels ...<-chan any) <-chan any {
	result := make(chan any)
	defer close(result)
	wg := sync.WaitGroup{}
	wg.Add(len(channels))
	for _, channel := range channels {
		go func(ch <-chan any, wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				select {
				case _, ok := <-ch:
					if !ok {
						return
					}
				}
			}
		}(channel, &wg)
	}
	wg.Wait()
	return result
}
