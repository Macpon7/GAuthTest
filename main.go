package main

import (
    "context"
    "encoding/base64"
    "fmt"
    "io/ioutil"
    "log"
    "crypto/rand"
    "net/http"
    "os"
    "time"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

func main() {


    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/", nilHandler)
    http.HandleFunc("/login", loginHandler)
    // http.HandleFunc("/logout", logoutHandler)
    http.HandleFunc("/oauth2callback", oathCallBackHandler)

    fmt.Println("Listening on port " + port)
    log.Fatal(http.ListenAndServe(":" + port, nil))
}

func nilHandler (w http.ResponseWriter, r *http.Request) {
    fmt.Println("This is the default handler")
}

var googleOauthConfig = &oauth2.Config{
    RedirectURL:    "https://goauthtest.herokuapp.com/oauth2callback",
    //ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
    ClientID:       "696387342528-o0tfdf782qltode6rlshaj5kgao3q75v.apps.googleusercontent.com",
    //ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
    ClientSecret:   "qcHUOA4u_FE-xDdewTRSRxVY",
    Scopes:         []string{"https://www.googleapis.com/auth/userinfo.email"},
    Endpoint:       google.Endpoint,
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func loginHandler(w http.ResponseWriter, r *http.Request) {

    // Create oauthState cookie
    oauthState := generateStateOauthCookie(w)
    u := googleOauthConfig.AuthCodeURL(oauthState)
    http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
    var expiration = time.Now().Add(365 * 24 * time.Hour)

    b := make([]byte, 16)
    rand.Read(b)
    state := base64.URLEncoding.EncodeToString(b)
    cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
    http.SetCookie(w, &cookie)

    return state
}

func oathCallBackHandler(w http.ResponseWriter, r *http.Request) {
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

func getUserDataFromGoogle(code string) ([]byte, error) {
    // Use code to get token and get user info from Google.

    token, err := googleOauthConfig.Exchange(context.Background(), code)
    if err != nil {
        return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
    }
    response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
    if err != nil {
        return nil, fmt.Errorf("failed getting user info: %s", err.Error())
    }
    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, fmt.Errorf("failed read response: %s", err.Error())
    }
    return contents, nil
}