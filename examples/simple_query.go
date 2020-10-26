package main

import (
	"fmt"
	"github.com/4thel00z/libdns"
	"github.com/4thel00z/libdns/utils"
)

func main() {
	response, err := libdns.SimpleQueryOnce(utils.CloudflarePrimary, "ransomware.host", utils.A, utils.InternetClass, 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", response)
}
