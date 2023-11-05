package main

import (
	"sync"
	"time"

	"github.com/Geun-Oh/initial-study/smtp"
)

var globalValue int

func action(i int, wg *sync.WaitGroup) {
	globalValue += i
	time.Sleep(1 * time.Second)

	wg.Done()
}

func main() {
	// startTime := time.Now()

	// var wg sync.WaitGroup
	// wg.Add(100)

	// for i := 0; i < 100; i++ {
	// 	go action(i, &wg)
	// }

	// wg.Wait()

	// delta := time.Now().Sub(startTime)
	// fmt.Printf("Result is %d, done in %.3fs.\n", globalValue, delta.Seconds())

	// err := rw.ReadAndWrite()
	// if err != nil {
	// panic(err)
	// }
	smtp.Smtp()
}
