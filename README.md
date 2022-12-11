# example-golang-scanner

### Description

A bare bones tcp port scanner; Autodiscover local network && scan neighbors for open ports


### Proof Of Concept Scripts

- [X] /hack/scan-ip.go: No concurrency; Single IP; Continous range of ports
- [X] /hack/print-popular-ports.go: Query && print out popular ports
- [X] /hack/scan-ip-top-ports.go: No concurrency; Single IP; Use the top X popular ports
- [X] /hack/scan-ip-tcp-concurrency.go: Port-scan concurrency; Single IP; Use the top X popular ports
- [X] /hack/scan-network.go: Port-scan concurrency accross a network; Use the top X popular ports
- [ ] /hack/scan-network-ip-concurrency.go: Port-scan && IP-scan concurrency across a network; Use the top X popular ports
- [ ] /hack/arp-scan.go: Discover active entities in the network

### Objectives

- [ ] Port-scan concurrency
- [ ] IP-scan concurrency
- [ ] Instead of continous range of ports, use the top X popular ports
- [ ] Unit Tests
- [ ] Some kind of E2E testing, potentially using another scanner to compare results 


### Thoughts && Conclusions

- Run a WaitGroup add *outside* of the goroutine, otherwise will race past your WaitGroup wait call
- Unconstrained goroutine execution leads to tons of open file descritors && slowed/stopped performance
- It seems like most network scanners will use ARP to discover what hosts are present && then scan those hosts; this seems more efficient
