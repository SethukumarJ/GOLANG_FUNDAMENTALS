package main

import (
	"fmt"
	"net"
)

func main() {

	ipRecords, _ := net.LookupIP("facebook.com")

	for _, ip := range ipRecords {

		fmt.Println(ip)
	}

}
