package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/pthomison/example-golang-scanner/poc/hack"
)

func main() {
	wg := new(sync.WaitGroup)

	start := time.Now()

	fmt.Printf("Scanning IP (%v): ", hack.IPAddr)

	for _, port := range hack.Ports {
		go func(port int) {
			wg.Add(1)
			if hack.IsPortOpen(hack.IPAddr, port) {
				fmt.Printf("%v ", port)
			}
			wg.Done()
		}(port)
	}

	wg.Wait()

	fmt.Println()

	hack.Elapsed(start)

}
