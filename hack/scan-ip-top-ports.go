package main

import (
	"fmt"

	"github.com/pthomison/example-golang-scanner/hack/common"
	"github.com/pthomison/example-golang-scanner/utils"
)

func main() {
	fmt.Printf("Scanning IP (%v): ", common.IPAddr)

	for _, port := range common.Ports {
		if utils.IsPortOpen(common.IPAddr, port) {
			fmt.Printf("%v ", port)
		}
	}

	fmt.Println()
}
