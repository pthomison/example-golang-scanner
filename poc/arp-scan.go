package main

import (
	"fmt"

	"github.com/pthomison/example-golang-scanner/poc/hack"
)

func main() {
	hack.NetworkARP(hack.IPNetwork, func(result *hack.ARPResult) {
		fmt.Printf("%v online: %v %v\n", result.IP, result.HWAddr, result.ResponseTime)
	})

}
