package os

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// 进度条,信号接收
func signal2() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)

	// Block until a signal is received.
	go func() {
		if _, ok := <-s; ok {
			fmt.Println("done")
			os.Exit(0)
		}
	}()

	d := time.Duration(time.Second)
	t := time.NewTicker(d)

	defer t.Stop()
	c := make(chan int)
	go func() {
		time.Sleep(time.Second * 10)
		c <- 1
	}()
L:
	for {
		select {
		case <-c:
			break L
		case <-t.C:
			fmt.Printf(".")
		}
	}
}
