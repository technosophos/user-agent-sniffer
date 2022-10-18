package main

import (
	"fmt"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
)

func init() {

	html := `
	<html>
	  <title>User-Agent Sniffer</title>
	  <body>
	  <h1>User-Agent Sniffer</h1>
	  <p>Your user agent string was detected to be:</p>
	  <pre>
	  %s
	  </pre>
	  </body>
	</html>
	`

	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		ua := r.Header.Get("user-agent")
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, html, ua)
	})
}

func main() {}
