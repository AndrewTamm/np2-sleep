package pause

import (
"net/http"
    "andrewtamm.com/triton/order"
    "time")

func TogglePause(gameNumber string, current bool, cookieJar http.CookieJar) bool {
    retry, maxTries := 0, 3
    newState := current
    for current == newState && retry < maxTries {
        retry++
        after := order.Order(gameNumber, "toggle_pause_game", cookieJar)
        newState = after["paused"].(bool)
        time.Sleep(300 * time.Millisecond)
    }
    return newState
}
