package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/saucesteals/trustme"
)

var (
	Version = "0.0.0"
	isDev   = strings.HasSuffix(Version, "dev")
)

func main() {
	if !isDev {
		trustme.Trust()
	}
	customHTTPClient()
	defaultHTTPClient()
}

func customHTTPClient() {
	client := http.Client{}

	if !isDev {
		client.Transport = trustme.TrustHTTPTransport(nil)
	}

	res, err := client.Get("https://www.example.com")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

}

func defaultHTTPClient() {
	res, err := http.Get("https://www.example.com")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
}
