package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	ips, err := net.LookupIP("DOMAINNAME")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not gather IPs: %v\n", err)
		os.Exit(1)
	}
	for _, ip := range ips {
		fmt.Printf("DOMAINNAME. IN A %s\n", ip.String())
	}
}

