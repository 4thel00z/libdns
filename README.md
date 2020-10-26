# libdns

![libdns-tests](https://github.com/4thel00z/libdns/workflows/Test/badge.svg)


Library to query DNS servers while packing the DNS query manually.

## Todo

* Add parsing the response portion ☑
* Add tests ☑
* Add dns record constants ☑
* Add question parsing  ☑
* Add high-level API (first high-level API added) ☐
* Add documentation ☐
* Add examples ☐

## Example usage

This demonstrates how you can make a simple DNS query. 

```go
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
```

## Acknowledgements

This project has drawn inspiration from [go-dnsquery](https://github.com/vishen/go-dnsquery/) and improves
on the sending portion.
