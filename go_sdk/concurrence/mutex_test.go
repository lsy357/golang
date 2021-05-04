package main

import (
	"sync"
	"testing"
	"time"
)

type mutexTestStruct struct {
	mutex   sync.RWMutex
	entries map[int]int
}

// 读写时都需要显示 lock，一个操作 lock，另一个不会被 block；同时读写会 panic
func TestUnlockRead(t *testing.T) {
	s := &mutexTestStruct{
		mutex:   sync.RWMutex{},
		entries: make(map[int]int),
	}

	var (
		len = 100000
		wg  = sync.WaitGroup{}
	)

	wg.Add(1)
	// 加锁耗时写
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)

		t.Logf("lock write, time: %v", time.Now().UnixNano())
		s.mutex.Lock()
		for i := 0; i < len; i++ {
			s.entries[i] = i + 1
		}
		t.Logf("release lock, time: %v", time.Now().UnixNano())
		s.mutex.Unlock()
	}()

	wg.Add(1)
	// 加锁写时读: 1.读操作 block 到写完后开始，读到数据为新数据，oldCount = 0; 2.读操作不 block，oldCount != 0
	// 结果：同时读写会 panic
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)

		oldCount := 0
		t.Logf("unlock read, time: %v", time.Now().UnixNano())
		for i := 0; i < len; i++ {
			v := s.entries[i]
			if v == i {
				oldCount++
			}
		}
		t.Logf("read done, old count: %v, time: %v", oldCount, time.Now().UnixNano())
	}()

	wg.Wait()
}
