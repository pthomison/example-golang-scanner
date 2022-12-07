package main

import (
	"fmt"

	"github.com/jreisinger/nmapservices"
	"github.com/pthomison/errcheck"
)

func main() {
	portcount := 100

	// Ports := utils.GetPopularPorts(1000)

	services, err := nmapservices.Get()
	errcheck.Check(err)
	topNmapPorts := services.Tcp().Top(portcount)

	for _, v := range topNmapPorts {
		fmt.Printf("%v ", v.Port)
	}

	fmt.Println()

	// spew.Dump(topNmapPorts)

}
