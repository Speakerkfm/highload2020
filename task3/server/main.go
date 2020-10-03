package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

const address = "localhost:8080"

func main() {
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("started listening: %s\n", address)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("fail to accept connection")
			continue
		}

		log.Printf("accepted new connection: %s\n", conn.RemoteAddr().String())

		go game(conn)
	}
}

func game(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	rand.Seed(time.Now().Unix())
	digit := rand.Int() % 100

loop:
	for {
		req, _, err := reader.ReadLine()
		if err != nil {
			log.Printf("fail to read request")
			if writeErr := writeResponse(writer, err.Error()); writeErr != nil {
				break
			}
			continue
		}

		log.Printf("got new request: %s\n", string(req))

		args := strings.Split(string(req), " ")
		if len(args) != 2 {
			if writeErr := writeResponse(writer, "Invalid command"); writeErr != nil {
				break
			}

			continue
		}

		switch args[0] {
		case "guess":
			userDigit, err := strconv.Atoi(args[1])
			if err != nil {
				if writeErr := writeResponse(writer, err.Error()); writeErr != nil {
					break
				}

				continue
			}

			if userDigit > digit {
				if err := writeResponse(writer, "more"); err != nil {
					break
				}

				continue
			}

			if userDigit < digit {
				if err := writeResponse(writer, "less"); err != nil {
					break
				}

				continue
			}

			if userDigit == digit {
				if err := writeResponse(writer, "correct"); err != nil {
					break
				}

				if err := writeResponse(writer, "You win!"); err != nil {
					break
				}

				break loop
			}
		default:
			if err := writeResponse(writer, "Invalid command"); err != nil {
				break
			}
			continue
		}
	}
}

func writeResponse(rw *bufio.Writer, msg string) error {
	if _, writeError := rw.WriteString(msg + "\n"); writeError != nil {
		log.Printf("fail to write response, cause: %s\n", writeError.Error())
		return writeError
	}

	if err := rw.Flush(); err != nil {
		log.Printf("fail to flush writer, cause: %s\n", err.Error())
		return err
	}

	return nil
}
