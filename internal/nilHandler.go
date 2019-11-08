package internal

import (
    "fmt"
    "net/http"
)

func NilHandler (w http.ResponseWriter, r *http.Request) {
    fmt.Println("This is the default handler")
}