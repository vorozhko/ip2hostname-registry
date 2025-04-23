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

/*
*
resolve takes an IP address as input and returns a slice of hostnames associated with it.
*
*/
func resolve(ip string) []string {
	if len(ipregistry[ip]) > 0 {
		return ipregistry[ip]
	}
	hostnames := search(ip)
	if hostnames == nil {
		fmt.Printf("No hostnames found for IP: %s\n", ip)
		return []string{}
	}
	ipregistry[ip] = append(ipregistry[ip], hostnames...)
	return ipregistry[ip]
}

/*
*
search performs a reverse DNS lookup
*
*/
func search(name string) []string {

	hostnames, err := net.LookupAddr(name)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil
	}
	return hostnames
}

func main() {
	ips := []string{"8.8.8.8", "104.21.56.30", "3.162.79.87", "1.1.1.1", "9.9.9.9"}
	results := make(chan string)
	for _, ip := range ips {
		go func(ip string) {
			hosts := resolve(ip)
			for _, h := range hosts {
				results <- fmt.Sprintf("%s = %s", ip, h)
			}
		}(ip)
	}

	for i := 0; i < len(ips); i++ {
		fmt.Println(<-results)
	}
}
