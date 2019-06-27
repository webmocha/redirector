package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	portEnv     = "PORT"
	defaultPort = "8080"
	targetEnv   = "TARGET"
	redirectEnv = "REDIRECT"
)

func main() {
	target := os.Getenv(targetEnv)
	redirect := http.StatusFound
	redirectStr := os.Getenv(redirectEnv)
	if redirectStr != "" {
		var redirectErr error
		redirect, redirectErr = strconv.Atoi(redirectStr)
		if redirectErr != nil {
			log.Fatalf("Error parsing REDIRECT: %+v\n", redirectErr)
		}
	}
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Error parsing TARGET %s: %+v\n", target, err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Scheme = targetURL.Scheme
		r.URL.Host = targetURL.Host
		http.Redirect(w, r, r.URL.String(), redirect)
	})

	port := os.Getenv(portEnv)
	if port == "" {
		port = defaultPort
	}

	fmt.Printf("Redirector listening on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
