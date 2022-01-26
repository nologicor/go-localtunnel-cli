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

	localAddress := os.Args[1]

	fmt.Println(localAddress)

	localPort, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	remoteHost := os.Args[3]

	subdomain := os.Args[4]

	s := strings.Split(remoteHost, "://")
	remoteAddrres := s[0] + "://" + subdomain + "." + s[1]

	log.Printf("Local address %s:%d tunnel to %s", localAddress, localPort, remoteAddrres)

	lt, err := localtunnel.New(
		localPort,
		localAddress,
		localtunnel.Options{
			Subdomain: subdomain,
			BaseURL:   remoteHost,
		},
	)
	if err != nil {
		log.Fatalf("localtunnel initialization error: %s", err.Error())
	}

	defer lt.Close()

	listener, err := localtunnel.Listen(localtunnel.Options{
		Subdomain: subdomain,
		BaseURL:   remoteHost,
	})
	if err != nil {
		log.Fatalf("listener error: %s", err.Error())
	}

	log.Println("localtunnel is running...")

	if err := http.Serve(listener, nil); err != nil {
		log.Fatalf("server error: %s", err.Error())
	}
}
