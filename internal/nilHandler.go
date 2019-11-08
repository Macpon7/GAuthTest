package internal

import (
    "fmt"
    "net/http"
)

func NilHandler (w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "This is the default handler. Try /login to login through Google and display your info here.")
}