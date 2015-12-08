package main

import (
	"net/http"
	"fmt"
	"os"
	"net"
)

var redirectHost string
var redirectHTTP bool

func main() {
	//Handle all requests
	http.HandleFunc("/",handle)
	//See if the there is a redirect host so we can decide later whether or not to use it
	redirectHost = os.Getenv("REDIRECT_HOSTNAME")

	//Check to see if HTTPS should be used.
	if os.Getenv("USE_HTTP") != "" {
		fmt.Println("Using https for redirects.")
		redirectHTTP = true;
	}
	fmt.Println("Listening on port 80...\n" + " Redirect hostname is " + os.Getenv("REDIRECT_HOSTNAME"))
	http.ListenAndServe(":80",nil)

}

func handle(w http.ResponseWriter, r * http.Request) {

	var host string

	//If the redirect host is empty use the host from the request
	if redirectHost == "" {
		host = net.SplitHostPort(r.Host)
	} else {
		host = redirectHost
	}

	fmt.Println("Processing request from " + r.RemoteAddr)
	path := r.RequestURI

	fmt.Printf("Redirecting to " + host+path)
	if !redirectHTTP {
		http.Redirect(w,r,"https://" + host + path,301)
	} else {
		http.Redirect(w,r,"http://" + host + path,301)
	}

}
