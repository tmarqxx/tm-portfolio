package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"tm-portfolio/handlers"
	"tm-portfolio/session"
)

const MAX_RETRIES = 10

func listen(port int) (net.Listener, error) {
	attempts := 0
	p := port

	for attempts < MAX_RETRIES {
		listener, err := net.Listen("tcp", ":"+strconv.Itoa(p))
		if err != nil {

			fmt.Printf("Port %d is already in use, trying again on port %d...\n", p, p+1)
			attempts++
			p++
			continue
		}

		fmt.Printf("Listening on port %d\n", p)
		return listener, nil
	}

	return nil, errors.New("No available ports between " + strconv.Itoa(port) + "-" + strconv.Itoa(port+MAX_RETRIES))
}

func main() {
	sh := session.NewMiddleware(handlers.NewDefaultMux())

	l, err := listen(3000)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.Serve(l, sh))
}
