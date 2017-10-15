package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Create a reverse proxy for the given target URI
func ReverseProxy(uri *url.URL, secretHeader string) *httputil.ReverseProxy {
	rp := httputil.NewSingleHostReverseProxy(uri)
	oldDirector := rp.Director
	rp.Director = func(r *http.Request) {
		oldDirector(r)
		r.Host = uri.Host
		r.Header.Del(secretHeader)
	}
	return rp
}

func main() {
	secret := flag.String("secret", "", "A secret to ensure the request is coming from the client")
	secretHeader := flag.String("secretHeader", "", "The header containing the client secret")
	target := flag.String("target", "", "A target uri to forward traffic to")
	flag.Parse()
	if *secret == "" || *secretHeader == "" || *target == "" {
		log.Fatal("Secret or target not provided!")
	}

	uri, err := url.Parse(*target)
	if err != nil {
		log.Fatal("Error parsing target uri", *target, err)
	}
	rp := ReverseProxy(uri, *secretHeader)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("New incoming request!")
		if r.Header.Get(*secretHeader) != *secret {
			log.Println("Incorrect or absent key!")
			return
		}
		log.Println("Correct key!")
		rp.ServeHTTP(w, r)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
