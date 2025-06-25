package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s domain\n", os.Args[0])
		fmt.Printf("%s requires a domain name\n", os.Args[0])
		os.Exit(1)
	}

	ips, err := net.LookupIP(os.Args[1])
	if err != nil {
		fmt.Printf("%s %s\n", os.Args[0], err)
		os.Exit(1)
	}

	if len(ips) == 0 {
		fmt.Printf("%s: failed to find the ip address for %s\n", os.Args[0], os.Args[1])
		os.Exit(1)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
