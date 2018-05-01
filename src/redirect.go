package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var redirectHost string
var redirectHTTP bool
var hostCheckEnabled bool
var whiteListedSuffix string
var redirectCode int
var redirectURL string

func main() {
	//Handle all requests
	http.HandleFunc("/", handle)
	//See if the there is a redirect host so we can decide later whether or not to use it
	redirectHost = os.Getenv("REDIRECT_HOSTNAME")
	whiteListedSuffix = os.Getenv("WHITELISTED_SUFFIX")
	redirectURL = os.Getenv("REDIRECT_URL")

	var err error
	if os.Getenv("REDIRECT_CODE") != "" {
		redirectCode, err = strconv.Atoi(os.Getenv("REDIRECT_CODE"))
	}
	if err != nil || redirectCode == 0 {
		redirectCode = 302
	}

	if redirectURL == "" && whiteListedSuffix != "" {
		log.Fatal("If you specify WHITELISTED_SUFFIX you must specify REDIRECT_URL!! Please Try Again!!")
	}

	if strings.HasPrefix(redirectURL, "https://") {
		redirectURL = strings.Replace(redirectURL, "https://", "", -1)
	} else if strings.HasPrefix(redirectURL, "http://") {
		redirectURL = strings.Replace(redirectURL, "http://", "", -1)
	}

	//Check to see if HTTPS should be used.
	if os.Getenv("USE_HTTP") != "" {
		fmt.Println("Using https for redirects.")
		redirectHTTP = true
	}

	if whiteListedSuffix == "" {
		if redirectHost == "" {
			fmt.Println("Listening on port 80...\n" + "Redirect hostname was not provided. Will use host from requests to redirect")
		} else {
			fmt.Println("Listening on port 80...\n" + "Redirect hostname is " + os.Getenv("REDIRECT_HOSTNAME") + ". Will use " + redirectHost + " for redirects.")
		}
	} else {
		hostCheckEnabled = true
		fmt.Println("Listening on port 80...\n" + "Whitelisted URL suffix is " + whiteListedSuffix + ". Will substitute " + redirectURL + " for un whitelisted domains.")
	}

	http.ListenAndServe(":80", nil)

}

func handle(w http.ResponseWriter, r *http.Request) {

	var host string

	host = r.Host

	//If the host contains a port then we want to remove it before continuing
	if strings.Contains(host, ":") {
		host = strings.Split(host, ":")[0]
	}

	// If the host check is enabled, go ahead and perform that function
	if hostCheckEnabled {
		if strings.HasSuffix(host, whiteListedSuffix) {
			host = r.Host
			redirect(w, r, host+r.RequestURI)
		} else {
			redirect(w, r, redirectURL)
		}
	} else {
		//If the redirect host is empty use the host from the request
		if redirectHost == "" {
			host = r.Host
		} else {
			host = redirectHost
		}
		redirect(w, r, host+r.RequestURI)
	}
}

func redirect(w http.ResponseWriter, r *http.Request, hostPath string) {
	fmt.Println("Processing request from " + r.RemoteAddr)
	path := r.RequestURI
	//If some passes http as a path then slap them on the hand with a bad request.
	if strings.HasPrefix(path, "/http:") || strings.HasPrefix(path, "/HTTP:") || strings.Contains(path, "comhttp") {
		fmt.Println("Someone is trying to do something nasty. Returning 400.")
		http.Error(w, "Bad Request", http.StatusBadRequest)
	} else {
		if !redirectHTTP {
			fmt.Println("Redirecting to https://" + hostPath)
			http.Redirect(w, r, "https://"+hostPath, redirectCode)
		} else {
			fmt.Println("Redirecting to http://" + hostPath)
			http.Redirect(w, r, "http://"+hostPath, redirectCode)
		}
	}

}
