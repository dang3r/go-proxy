# Go-Proxy

GoProxy is a simple proxy server written in Golang. Given a target url, it
proxies all requests to that target. A client header and secret are also used to
ensure only allowed clients can make proxy requests.

## Installation

To install go-proxy, use

```sh
go get github.com/dang3r/go-proxy
```

## Usage

```sh
$ go build proxy.go 
$ ./proxy -help
Usage of ./proxy:
  -secret string
        A secret to ensure the request is coming from the client
  -secretHeader string
        The header containing the client secret
  -target string
        A target uri to forward traffic to
```
