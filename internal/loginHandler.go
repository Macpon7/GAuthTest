package internal

import (
    "encoding/json"
    "fmt"
    "github.com/pkg/errors"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "io/ioutil"
    "net/http"
)

type gAuthInfo struct {
    RedirectUrl     string      `json:"RedirectUrl"`
    ClientID        string      `json:"ClientID"`
    ClientSecret    string      `json:"ClientSecret"`
}

var gAuthConfig gAuthInfo

var googleOauthConfig = &oauth2.Config{
    Scopes:         []string{"https://www.googleapis.com/auth/userinfo.email"},
    Endpoint:       google.Endpoint,
}

func AuthConfigInit() error {
    file, err := ioutil.ReadFile("GAuthInfo.json")
    if err != nil {
        fmt.Println(err.Error())
        return errors.New("Could not find file GAuthConfig")
    }

    err = json.Unmarshal([]byte(file), &gAuthConfig)
    if err != nil {
        fmt.Println(err.Error())
        return errors.New("Could not decode GAuthConfig file")
    }

    googleOauthConfig.RedirectURL = gAuthConfig.RedirectUrl
    googleOauthConfig.ClientID = gAuthConfig.ClientID
    googleOauthConfig.ClientSecret = gAuthConfig.ClientSecret

    return nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

    // Create oauthState cookie
    oauthState := generateStateOauthCookie(w)
    u := googleOauthConfig.AuthCodeURL(oauthState)
    http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}