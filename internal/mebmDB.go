package internal

import (
    "fmt"
    "net/http"
)

type userInfo struct {
    Email       string      `json:"email"`
}

var users []userInfo

func dbSave(in userInfo) {
    users = append(users, in)
}

func dbPrintAll(w http.ResponseWriter) {
    fmt.Fprint(w, users)
}