poc_one:
	go run ./poc/scan-ip.go

print_ports:
	go run ./poc/print-popular-ports.go

poc_two:
	go run ./poc/scan-ip-top-ports.go

poc_three:
	go run ./poc/scan-ip-tcp-concurrency.go

poc_four:
	go run ./poc/scan-network.go

poc_five:
	go run ./poc/scan-network-concurrency.go