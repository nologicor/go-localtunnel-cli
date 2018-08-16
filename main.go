package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/jonasfj/go-localtunnel"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "localtunnel <local address> <local port> <remote host:port> <subdomain>")
		os.Exit(1)
	}

	laddr := os.Args[1]

	lport, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	rhost := os.Args[3]

	subdomain := os.Args[4]

	s := strings.Split(rhost, "://")
	raddr := s[0] + "://" + subdomain + "." + s[1]

	log.Printf("Local address %s:%d tunnel to %s", laddr, lport, raddr)

	_, err = localtunnel.New(lport, laddr, localtunnel.Options{Subdomain: subdomain, BaseURL: rhost})
	if err != nil {
		panic(err)
	}
	listener, err := localtunnel.Listen(localtunnel.Options{})
	if err != nil {
		panic(err)
	}
	log.Println("localtunnel is running...")
	http.Serve(listener, nil)
}
