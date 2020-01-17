package os2

import (
	"net/http"
	"sync"
)

func waitgroup() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			http.Get("https://httpstat.us/200")
			wg.Done()
		}()
	}
	wg.Wait()
}
