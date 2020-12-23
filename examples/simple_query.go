package main

import (
	"fmt"
	"github.com/4thel00z/libdns/v1"
	"github.com/4thel00z/libdns/v1/pkg/libdns/utils"
)

func main() {
	response, err := v1.SimpleQueryOnce(utils.CloudflarePrimary, "ransomware.host", utils.A, utils.InternetClass, 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", response)
}
