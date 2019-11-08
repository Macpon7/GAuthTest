package cmd

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "GAuthTest/internal"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/", internal.NilHandler)
    http.HandleFunc("/login", internal.LoginHandler)
    // http.HandleFunc("/logout", logoutHandler)
    http.HandleFunc("/oauth2callback", internal.OathCallBackHandler)

    fmt.Println("Listening on port " + port)
    log.Fatal(http.ListenAndServe(":" + port, nil))
}