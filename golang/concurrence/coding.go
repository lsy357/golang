package main

import (
	"context"
	"fmt"
)

// 4个协程，1234号协程只能打印1234数字，打印1234 2341 3412 4123
func main() {
	targetSeq := "1234 2341 3412 4123"
	chars := []rune(targetSeq)
	channels := make([]chan byte, 4)
	done := make(chan byte)
	ctx, cancel := context.WithCancel(context.Background())

	for i := range channels {
		channels[i] = make(chan byte)
		go func(ctx context.Context, num int) {
			for {
				select {
				case <-channels[num]:
					fmt.Printf("%d", num+1)
					done <- 1
				case <-ctx.Done():
					break
				}
			}
		}(ctx, i)
	}
	for _, char := range chars {
		if char < '1' || char > '4' {
			continue
		}
		channels[char-'1'] <- 1
		<-done
	}
	cancel()
	for i := range channels {
		close(channels[i])
	}
}
