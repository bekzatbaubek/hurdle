# Hurdle

Wordle clone in Go

## Run the game

```sh
git checkout task-1
go run ./main.go
```

## Code design decisions

- Added unit tests for Hurdle core logic

```sh
go test .
```

- Wordlist and number of rounds are loaded from Environment Variables


## Task 2 Design

- Client code in Typescript (vite)
- Server code in Go

- Game state is tracked with `SessionTokens` stored in HTTP Cookies (Prevent JS access)
- Server tracks each client session and their game state
- CORS Headers to allow client to access the server
