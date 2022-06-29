package trustme_test

import (
	"log"
	"net/http"

	"github.com/saucesteals/trustme"
)

func ExampleTrustTLSConfig() {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: trustme.TrustTLSConfig(nil),
		},
	}

	res, err := client.Get("https://www.example.com")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
}

func ExampleTrustHTTPTransport() {
	client := http.Client{
		Transport: trustme.TrustHTTPTransport(nil),
	}

	res, err := client.Get("https://www.example.com")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
}
