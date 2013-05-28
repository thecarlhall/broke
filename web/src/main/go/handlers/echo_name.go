package broke.handlers

import (
  "fmt"
  "net/http"
)

struct EchoName {
}

func (en *EchoName) handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
