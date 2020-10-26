package main

import (
	"fmt"
	l "github.com/4thel00z/libdns"
	u "github.com/4thel00z/libdns/utils"
)

func main() {
	response, err := l.SimpleQueryOnce(u.CloudflarePrimary, "ransomware.host", u.A, u.InternetClass, 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", response)
}
