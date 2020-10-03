package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const address = "localhost:8080"

func main() {
	conn, _ := net.Dial("tcp", address)
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text + "\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(message)
	}
}
