# libdns

![libdns-tests](https://github.com/4thel00z/libdns/workflows/Test/badge.svg)
![libdns-logo](https://github.com/4thel00z/libdns/raw/assets/logo.png)


Go DNS Library that exposes various DNS primitives to the user such as:

- DNS Servers
- Models for the DNS format
- Various DNS constants
- etc.

Besides that it sports a high-level API to enable simple programs that want to query DNS servers and not bother with the nitty-gritty.

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

## Todo

* Add parsing the response portion ☑
* Add tests ☑
* Add dns record constants ☑
* Add question parsing  ☑
* Add high-level API (first high-level API added) ☐
* Add documentation ☐
* Add examples ☐

## Acknowledgements

This project has drawn inspiration from [go-dnsquery](https://github.com/vishen/go-dnsquery/).

## License

This project is licensed under the GPL-3 license.
