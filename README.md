[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/thomasdseao/go-graylog) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/thomasdseao/go-graylog/master/LICENSE) [![CodeFactor](https://www.codefactor.io/repository/github/thomasdseao/go-graylog/badge)](https://www.codefactor.io/repository/github/thomasdseao/go-graylog)
<h1 align="center">Go Graylog Package</h1>

### This package implement Graylog GELF interface to send message using UDP or TCP transport, in Golang.


## Example

```go
package main

import (
	"github.com/thomasdseao/go-graylog"
)

func main() {
	// Create Gelf instance
	gelf := graylog.NewGelf(graylog.Config{
		"graylog1.example.com",
		2202,
		graylog.UDP,
		true,
	})
    
	// Create message and JSON encode it
	message := graylog.Message{
		Version:      "1.1",
		Host:         "example.com",
		ShortMessage: "This is the short message",
	}
	jsonMessage, _ := json.Marshal(message)
	
	// Send message
	sent, err := gelf.Send(jsonMessage)

}
```

## Tests
```
go test
```

## Contribute
#### Feel free to submit PR with another implementations of Graylog or improvements.
