# example-golang-scanner

### Description

A bare bones tcp port scanner; Autodiscover local network && scan neighbors for open ports


### Proof Of Concept Scripts

- [X] /poc/scan-ip.go: No concurrency; Single IP; Continous range of ports
- [X] /poc/print-popular-ports.go: Query && print out popular ports
- [X] /poc/scan-ip-top-ports.go: No concurrency; Single IP; Use the top X popular ports
- [X] /poc/scan-ip-tcp-concurrency.go: Port-scan concurrency; Single IP; Use the top X popular ports
- [X] /poc/scan-network.go: Port-scan concurrency accross a network; Use the top X popular ports
- [X] /poc/arp-scan.go: Discover active entities in the network
- [ ] /poc/scan-network-arp-discovery.go: Port-scan && IP-scan concurrency across a network; Use the top X popular ports


### Objectives

- [ ] Port-scan concurrency
- [ ] IP-scan concurrency
- [ ] Instead of continous range of ports, use the top X popular ports
- [ ] Unit Tests
- [ ] Some kind of E2E testing, potentially using another scanner to compare results 


### TODOs

- [ ] 


### Thoughts && Conclusions

- Run a WaitGroup add *outside* of the goroutine, otherwise will race past your WaitGroup wait call
- Unconstrained goroutine execution leads to tons of open file descritors && slowed/stopped performance
- It seems like most network scanners will use ARP to discover what hosts are present && then scan those hosts; this seems more efficient
- Figuring the correct wait time for the TCP scan is key (potentially derive from ARP or ICMP ping time)