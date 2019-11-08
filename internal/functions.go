package internal

import (
    "context"
    "crypto/rand"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/pkg/errors"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="


func generateStateOauthCookie(w http.ResponseWriter) string {
    var expiration = time.Now().Add(365 * 24 * time.Hour)

    b := make([]byte, 16)
    rand.Read(b)
    state := base64.URLEncoding.EncodeToString(b)
    cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
    http.SetCookie(w, &cookie)

    return state
}

func getUserDataFromGoogle(code string) (userInfo, error) {
    // Use code to get token and get user info from Google.
    var tempUser userInfo

    token, err := googleOauthConfig.Exchange(context.Background(), code)
    if err != nil {
        fmt.Println("Code exchange failed: " + err.Error())
        return tempUser, errors.New("Code exchange failed")
    }
    response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
    if err != nil {
        fmt.Println("Failed getting user info: " + err.Error())
        return tempUser, errors.New("Failed getting user info")
    }
    defer response.Body.Close()

    err = json.NewDecoder(response.Body).Decode(&tempUser)
    if err != nil {
        fmt.Println("Failed decoding user info from google: " + err.Error())
        return tempUser, errors.New("Failed decoding user info from google")
    }

    return tempUser, nil
}