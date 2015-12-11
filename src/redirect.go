package main

import (
	"net/http"
	"fmt"
	"os"
)

var redirectHost string
var redirectHTTP bool

func main() {
	//Handle all requests
	http.HandleFunc("/", handle)
	//See if the there is a redirect host so we can decide later whether or not to use it
	redirectHost = os.Getenv("REDIRECT_HOSTNAME")

	//Check to see if HTTPS should be used.
	if os.Getenv("USE_HTTP") != "" {
		fmt.Println("Using https for redirects.")
		redirectHTTP = true;
	}

	if redirectHost == "" {
		fmt.Println("Listening on port 80...\n" + "Redirect hostname was not provided. Will use host from requests to redirect")
	} else {
		fmt.Println("Listening on port 80...\n" + "Redirect hostname is " + os.Getenv("REDIRECT_HOSTNAME") + ". Will use " + redirectHost + " for redirects.")
	}

	http.ListenAndServe(":80", nil)

}

func handle(w http.ResponseWriter, r *http.Request) {

	var host, port string
	//If the redirect host is empty use the host from the request
	if redirectHost == "" {
		host = r.Host
	} else {
		host = redirectHost
	}

	fmt.Println("Processing request from " + r.RemoteAddr)
	path := r.RequestURI


	if !redirectHTTP {
		fmt.Println("Redirecting to https://" + host + path + " ignoring port " + port)
		http.Redirect(w, r, "https://" + host + path, 301)
	} else {
		fmt.Println("Redirecting to http://" + host + path + " ignoring port " + port)
		http.Redirect(w, r, "http://" + host + path, 301)
	}

}
