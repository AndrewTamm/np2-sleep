package main

import (
    "fmt"
    "andrewtamm.com/triton/login"
    "net/http/cookiejar"
    "andrewtamm.com/triton/order"
    "andrewtamm.com/triton/pause"
    "flag"
    "os"
    "andrewtamm.com/triton/gmail")

var user, pass, game, to, from, token string
var sendEmail bool

func main() {
    cookieJar, _ := cookiejar.New(nil)

    login.SignIn(user, pass, cookieJar)
    before := order.Order(game , "full_universe_report", cookieJar)["paused"].(bool)
    after := pause.TogglePause(game, before, cookieJar)
    if sendEmail {
        sendStatusUpdateMail(after)
    }
}

func sendStatusUpdateMail(isPaused bool) {
    var message string
    if isPaused {
        message = fmt.Sprintf("Game ID: %s is now paused.", game)
    } else {
        message = fmt.Sprintf("Game ID: %s is now running.", game)
    }

    gmail.Send(to, from, fmt.Sprintf("Neptune's Pride II Game %s: status update", game), message, token)
}

func init() {
    flag.StringVar(&user, "login", "Required", "Game administrator login")
    flag.StringVar(&user, "l", "Required", "Game administrator login (shorthand)")
    flag.StringVar(&pass, "password", "Required", "Account Password")
    flag.StringVar(&pass, "p", "Required", "Account Password (shorthand)")
    flag.StringVar(&game, "gameId", "Required", "Target Game ID, must be administrator")
    flag.StringVar(&game, "g", "Required", "Target Game ID, must be administrator (shorthand)")
    flag.StringVar(&to, "to", "Optional", "Email address to send status notifications")
    flag.StringVar(&from, "from", "Optional", "Sending email address")
    flag.StringVar(&token, "token", "Optional", "Gmail application token")
    flag.Parse()

    for _, value := range [...]string{user, pass, game} {
        checkRequiredStringVar(value)
    }

    checkEmailParametersPresent()
}

func checkRequiredStringVar(variable string) {
    if (variable == "Required") {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
        flag.PrintDefaults()
        os.Exit(1)
    }
}

func checkEmailParametersPresent() {
    var anyPresent, allPresent = false, true

    for _, value := range [...]string{to, from, token} {
        anyPresent = anyPresent || value != "Optional"
        allPresent = allPresent && value != "Optional"
    }

    if anyPresent && !allPresent {
        fmt.Fprint(os.Stderr, "All email parameters are required if any are supplied")
        os.Exit(1)
    }

    sendEmail = allPresent
}
