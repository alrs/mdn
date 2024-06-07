package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/eternal-flame-AD/go-termux"
	"github.com/pd0mz/go-maidenhead"
)

const timeout = time.Duration(10 * time.Second)

func main() {
	gpsCTX, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	mode := termux.GPS
	l, err := termux.Location(gpsCTX, mode)
	if err != nil {
		log.Printf("ignored error on GPS location: %q", err)
		netCTX, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		mode = termux.Network
		l, err = termux.Location(netCTX, mode)
		if err != nil {
			log.Fatalf("error on network location: %q", err)
		}
	}
	log.Printf("successful location mode: %s", mode)
	p := maidenhead.Point{l.Latitude, l.Longitude}
	m, err := p.Locator(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
	// err = termux.Toast(m, termux.ToastOption{})
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
}
