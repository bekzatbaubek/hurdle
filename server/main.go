package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"hurdle-server/hurdle"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type GameInitResponse struct {
	PlayerToken string `json:"player_token"`
}

type GameGuessRequest struct {
	Guess string `json:"guess"`
}

type Guess struct {
	Round int    `json:"round"`
	Guess string `json:"guess"`
	Hint  string `json:"hint"`
}

type GameSession struct {
	CurrentRound int     `json:"current_round"`
	Guesses      []Guess `json:"guesses"`
}

func init() {
	os.Setenv("CANDIDATES", "HELLO,WORLD,QUITE,FANCY,FRESH,PANIC,CRAZY,BUGGY,SCARE")
	os.Setenv("NUMBER_OF_ROUNDS", "5")
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	candidate_env := os.Getenv("CANDIDATES")
	candidates := strings.Split(candidate_env, ",")

	numberOfRoundsEnv := os.Getenv("NUMBER_OF_ROUNDS")
	numberOfRounds, err := strconv.Atoi(numberOfRoundsEnv)
	if err != nil {
		// Configuration error
		panic(err)
	}

	answer := hurdle.PickAnswer(candidates)

	log.Printf("Hurdle configuration loaded, number of rounds = %d, answer = %s", numberOfRounds, answer)

	gameSessions := make(map[string]*GameSession)

	// 1. Init server
	// 2. Start server
	// 3. Handle requests
	// 3.1 Generate session tokens for new players (for tracking progress)
	// 3.2 Handle player guesses

	http.HandleFunc("/api/sessions", func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("SessionToken")
		if err != nil {
			buffer := make([]byte, 16)

			_, err := rand.Read(buffer)
			if err != nil {
				log.Printf("error when generating player token = %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}

			response := GameInitResponse{
				PlayerToken: hex.EncodeToString(buffer),
			}
			log.Printf("Generated player token = %s", response.PlayerToken)
			gameSessions[response.PlayerToken] = &wGameSession{
				CurrentRound: 1,
				Guesses:      []Guess{},
			}

			marshalled, err := json.Marshal(response)
			if err != nil {
				log.Printf("error when marshalling response to user = %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			http.SetCookie(w, &http.Cookie{
				Name:     "SessionToken",
				Value:    response.PlayerToken,
				Path:     "/",
				MaxAge:   3600,
				Secure:   false,
				SameSite: http.SameSiteLaxMode,
			})
			w.Header().Add("Access-Control-Allow-Origin", "http://localhost:5173")
			w.Header().Add("Access-Control-Allow-Credentials", "true")
			w.Write(marshalled)
		} else {
			log.Printf("Player token found = %s", cookie.Value)

			game, err := json.Marshal(gameSessions[cookie.Value])
			if err != nil {
				log.Printf("error when marshalling game session for player token = %s", cookie.Value)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}

			w.Header().Add("Access-Control-Allow-Origin", "http://localhost:5173")
			w.Header().Add("Access-Control-Allow-Credentials", "true")
			w.Write(game)
		}
	})

	http.HandleFunc("/api/guesses", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			cookie, err := r.Cookie("SessionToken")
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}

			var guessRequest GameGuessRequest
			err = json.NewDecoder(r.Body).Decode(&guessRequest)

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

			log.Printf("got guess = %s, from player = %s", guessRequest.Guess, cookie.Value)

			w.Header().Add("Access-Control-Allow-Origin", "http://localhost:5173")
			w.Header().Add("Access-Control-Allow-Credentials", "true")

			hint := hurdle.Hurdle(answer, guessRequest.Guess)

			gameSession := gameSessions[cookie.Value]

			gameSession.Guesses = append(gameSession.Guesses, Guess{
				Round: gameSession.CurrentRound,
				Guess: guessRequest.Guess,
				Hint:  hint,
			})

			gameSession.CurrentRound++

			log.Printf("game session state = %v", gameSession)

			response, err := json.Marshal(gameSession)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}

			w.Write(response)

		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Printf("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
