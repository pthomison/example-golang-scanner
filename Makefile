poc_one:
	go run ./hack/scan-ip.go

print_ports:
	go run ./hack/print-popular-ports.go

poc_two:
	go run ./hack/scan-ip-top-ports.go

poc_three:
	go run ./hack/scan-ip-tcp-concurrency.go

poc_four:
	go run ./hack/scan-network.go

poc_five:
	go run ./hack/scan-network-concurrency.go