# np2-sleep
Sleep timer for Neptune's Pride 2:Triton

## Index
1. [Goal](#goal)
2. [Usage](#usage)
 * [Basic Toggle](#basic-toggle)
 * [Sending update emails](#sending-update-emails)
 * [Setting up a cron](#setting-up-a-cron)

## Goal

Since Neptune's Pride runs around the clock, important events can happen at times when some or all of the players can't get to the game. Whatever your personal feelings on this, some don't like this aspect of the game and so this program very simply will toggle the pause status on the game.

## Usage

### Basic toggle

#### As a script
There's no need to directly build an executable;the program can be run directly as from the command line

`go run src/main.go -g <game-id> -l <admin-login> -p <admin-password>`

#### As an executable
Alternatively, compile and run it as an executable:

```sh
$ go build -o np2-toggle src/main.go
$ go run src/main.go -g <game-id> -l <admin-login> -p <admin-password>
```

### Sending update emails
There's little (read: no) error handling here, so it's technically possible than the toggle could fail and leave the game running past the desired cutoff. To make this easier to detect, you can send an email through gmail to notify the admin of updates: a missing email means something went wrong.

You can get a application specific password to use as the token from gmail itself, an exercise left up to the reader.

```sh
$ go run src/main.go -g <game-id> -l <admin-login> -p <admin-password> \
    -to <update-email> -from <admin-email> -token <gmail application token>
```

### Setting up a cron
This program is designed to be run as a scheduled cron job. 

Examples:

```sh
#!/bin/sh

export GOROOT=/usr/lib/go
export GOPATH=/home/ubuntu/triton

go run /home/ubuntu/triton/src/main.go $@
```

Then edit your crontab to run the toggle operation at the desired times (in this case, 10 am and 10pm PST when the system clock is set to UTC)

```bash
0 5 * * * /home/ubuntu/triton/triton.sh -g=<game-id> -l=<admin> -p=<password>
0 17 * * * /home/ubuntu/triton/triton.sh -g=<game-id> -l=<admin> -p=<password>
```

If you decide to go the route of building a binary, the shell script is unnecessary and you can run the program from the crontab directly. 
