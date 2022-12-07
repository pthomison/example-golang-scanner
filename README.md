# example-golang-scanner

### Description

A bare bones tcp port scanner; Autodiscover local network && scan neighbors for open ports


### Proof Of Concept Scripts

- [X] /hack/scan-ip.go: No concurrency; Single IP; Continous range of ports
- [X] /hack/scan-ip-top-ports.go: No concurrency; Single IP; Use the top X popular ports
- [X] /hack/scan-ip-tcp-concurrency.go: Port-scan concurrency; Single IP; Use the top X popular ports
- [X] /hack/scan-network.go: Port-scan concurrency accross a network; Use the top X popular ports
- [ ] /hack/scan-network-ip-concurrency.go: Port-scan && IP-scan concurrency across a network; Use the top X popular ports

### Objectives

- [ ] Port-scan concurrency
- [ ] IP-scan concurrency
- [ ] Instead of continous range of ports, use the top X popular ports


### Thoughts && Conclusions