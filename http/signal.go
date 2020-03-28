package http

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	done := make(chan bool, 1)
	s1 := make(chan os.Signal, 1)
	signal.Notify(s1, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-s1
		fmt.Println("signal sigint or sigterm")
		done <- true
	}()

	s2 := make(chan os.Signal, 1)
	signal.Notify(s2, syscall.SIGWINCH)

	go func() {
		for {
			<-s2
			fmt.Println("signal sigwinch")
		}
	}()
	<-done

	s3 := make(chan os.Signal, 1)
	signal.Notify(s3, syscall.SIGWINCH)
	signal.Ignore(syscall.SIGINT)
	go func() {
		<-s3
		fmt.Println("s3")
		signal.Stop(s3)
		<-s3
		done <- true
	}()
	<-done
}
