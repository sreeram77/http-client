# http-client

HTTP client written in go for sending requests. http-client is a wrapper on top of http.Client, providing an easy to use HTTP client.

## Installing
```sh
go get github.com/sreeram77/http-client
```

## Example
```go
package main

import (
	"fmt"
	"github.com/sreeram77/http-client"
)

func main() {
  // Create an instance of client
	c := http.New()
  // Make a GET request to URL
	r, err := c.Get("https://google.com", nil)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(r.StatusCode)
}
```
