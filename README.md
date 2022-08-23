<p align="center">
  <a href="https://travis-ci.org/thomasdseao/go-graylog" target="_blank"><img src="https://travis-ci.org/thomasdseao/go-graylog.png?branch=master"></a>
</p>

# Go Graylog package

### This package implement Graylog GELF interface to send message using UDP or TCP transport, in Golang.


# Example

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
    
	// Send message
    sent, err := gelf.Send(graylog.Message{
        Version:      "1.1",
        Host:         "example.com",
        ShortMessage: "This is the short message",
    })
	
}
```

# Tests
```
go test
```

# Contribute
#### Feel free to submit PR with another implementations of Graylog or improvements.