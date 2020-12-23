package main

import (
	"fmt"
	u "github.com/4thel00z/libdns/v1/pkg/libdns/utils"
	"github.com/4thel00z/libdns/v1"
)

func main() {
	response, err := v1.SimpleQueryOnce(u.CloudflarePrimary, "ransomware.host", u.A, u.InternetClass, 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", response)
}
