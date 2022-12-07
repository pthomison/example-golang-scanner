poc_one:
	go run ./hack/scan-ip.go

poc_two:
	go run ./hack/scan-ip-top-ports.go

poc_three:
	go run ./hack/scan-ip-tcp-concurrency.go