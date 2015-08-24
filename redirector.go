// Author: Patrick Jackson (@patjack)

package main

import (
    "fmt"
    "net/http"
    "strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.Host, "www") {
		fmt.Fprintf(w, "1") // shouldn't be hit directly, or show error message
    } else {
    	http.Redirect(w, r, fmt.Sprintf("http://www.%s%s", r.Host, r.URL.String()), 302)
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":80", nil)
}