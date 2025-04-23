package main

import (
	"fmt"
	"net"
)

var ipregistry = map[string][]string{
	"8.8.8.8":      {"dns.google"},
	"104.21.56.30": {"vorozhko.net"},
	"10.0.0.1":     {"dev-host1", "dev-host2"},
}

func resolve(ip string) []string {
	if len(ipregistry[ip]) > 0 {
		return ipregistry[ip]
	}
	hostnames := search(ip)
	if len(hostnames) > 0 {
		for _, hostname := range hostnames {
			ipregistry[ip] = append(ipregistry[ip], hostname)
		}
	}

	return ipregistry[ip]
}

func search(name string) []string {

	hostnames, err := net.LookupAddr(name)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil
	}
	return hostnames
}

func main() {
	for _, ip :=  range []string{"8.8.8.8", "104.21.56.30", "3.162.79.87", "1.1.1.1", "9.9.9.9"} {
		hosts := resolve(ip)
		for _, h := range hosts {
			fmt.Println(ip,"=",h)
		}	
	}
}
