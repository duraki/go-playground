package main

import (
	"net/url"
	"fmt"
	"net"
)

const URL = "redis://golang:parsing@example.com:1000/api"

func main() {

	u, err := url.Parse(URL)
	if err != nil {
		panic(err)
	}

	fmt.Println("URL Scheme:", u.Scheme)
	fmt.Println("URL Auth combo:", u.User)
	fmt.Println("URL Username:", u.User.Username())

	pass, _ := u.User.Password()
	fmt.Println("URL Password:", pass)

	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("URL host", host)
	fmt.Println("URL port", port)

	fmt.Println("URL path", u.Path)

}
