package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

// To simulate Mutex using Channels
func main() {
	runtime.GOMAXPROCS(4)

	// Previous logs will be erased due to new creations
	f, _ := os.Create("./log.txt")
	f.Close()

	// To give a large, nice buffer to try and prevent our application
	// from becoming IO bound,
	// since writing to the disk is going to be much slower than our calculation
	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				f, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
				if err != nil {
					fmt.Println(err)
				}
				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)
				f.Close()
			} else {
				break
			}
		}
	}()

	mutex := make(chan bool, 1)

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex <- true
			go func() {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
				<-mutex
			}()
		}
	}
}
