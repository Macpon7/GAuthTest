package internal

import (
    "fmt"
    "log"
    "net/http"
)

func OathCallBackHandler(w http.ResponseWriter, r *http.Request) {
    // Read oauthState from Cookie
    oauthState, _ := r.Cookie("oauthstate")

    if r.FormValue("state") != oauthState.Value {
        log.Println("invalid oauth google state")
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    data, err := getUserDataFromGoogle(r.FormValue("code"))
    if err != nil {
        log.Println(err.Error())
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    // GetOrCreate User in your db.
    // Redirect or response with a token.
    // More code .....
    fmt.Fprintf(w, "UserInfo: %s\n", data)
}