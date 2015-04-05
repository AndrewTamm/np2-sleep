package login

import (
    "net/http"
    "net/url"
    "fmt")

func SignIn(user, password string, cookieJar http.CookieJar) {
    client := &http.Client{
        Jar: cookieJar,
    }

    resp, err := client.PostForm("http://triton.ironhelmet.com/arequest/login",
    url.Values{"type": {"login"}, "alias": {user}, "password": {password}})

    fmt.Println(resp)

    if (err != nil) {
        fmt.Println(resp)
        panic(resp)
    }
}
