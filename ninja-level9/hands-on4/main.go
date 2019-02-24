package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			m.Lock()
			v := counter
			v++
			counter = v
			fmt.Println(counter)
			wg.Done()
			m.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
