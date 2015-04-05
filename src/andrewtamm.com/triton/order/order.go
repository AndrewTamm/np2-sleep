package order
import (
    "net/http"
    "net/url"
    "fmt"
    "io/ioutil"
    "encoding/json")

func Order(gameNumber, order string, cookieJar http.CookieJar) map[string]interface {} {
    client := &http.Client{
        Jar: cookieJar,
    }

    resp, err := client.PostForm("http://triton.ironhelmet.com/grequest/order",
    url.Values{"type": {"order"}, "version": {"7"}, "order": {order}, "game_number": {gameNumber}})

    if (err != nil) {
        fmt.Println(resp)
        panic(resp)
    }

    defer resp.Body.Close()

    return bodyToJson(resp)
}

func bodyToJson(resp *http.Response) map[string]interface {} {
    body, _ :=  ioutil.ReadAll(resp.Body)

    var f interface{}
    json.Unmarshal(body, &f)

    m := f.(map[string]interface{})
    return m["report"].(map[string]interface{})
}
