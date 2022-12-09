package main

import (
	"fmt"
	"sync"

	"github.com/pthomison/example-golang-scanner/hack/common"
	"github.com/pthomison/example-golang-scanner/utils"
)

var (
	wg = new(sync.WaitGroup)
)

func main() {
	fmt.Printf("Scanning IP (%v): ", common.IPAddr)

	for _, port := range common.Ports {
		go func(port int) {
			wg.Add(1)
			if utils.IsPortOpen(common.IPAddr, port) {
				fmt.Printf("%v ", port)
			}
			wg.Done()
		}(port)
	}

	wg.Wait()

	fmt.Println()
}
