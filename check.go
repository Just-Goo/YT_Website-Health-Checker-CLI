package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination, port string) string {
	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	
	conn, err := net.DialTimeout("tcp", address, timeout)

	if err != nil {
		return fmt.Sprintf("[DOWN] %v is unreachable, \n Error: %v", destination, err)	
	} 
	
	return fmt.Sprintf("[UP] %v is reachable, \n From: %v\n To: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
}
