package main

import (
  "log"
	"os"
	"github.com/joho/godotenv"
  "net/http"
  "fmt"
)

func main() {

    // Load the .env file and get the API_KEY
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    apiKey := os.Getenv("API_KEY")
    println(apiKey)

    // --USER ROUTES-- //
 
    // GAMES ROUTES //
    // GET ALL games 
    http.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "GET ALL games")
    })
    // GET a game by id 
    http.HandleFunc("/games/:id", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "GET a game by id")
    })
    // POST add a new game
    http.HandleFunc("/games/add", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "POST add a new game")
    })


    // CHARACTERS ROUTES //
    // GET all characters
    http.HandleFunc("/characters", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "GET all characters")
    })
    // GET a character by id 
    http.HandleFunc("/characters/:id", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "GET a character by id")
    })
    // POST add a new character 
    http.HandleFunc("/characters/add", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "POST add a new character")
    })
    

    // --OPENAI ROUTES-- //
    
    // GET the infos about a character
    http.HandleFunc("/openai/character/:game/:name", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "GET the infos about a character with his game and name")
    })
    
    // POST send a message to the chatbot and get the response
    http.HandleFunc("/openai/chatbot", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "POST send a message to the chatbot and get the response")
    })

	  http.ListenAndServe(":8080", nil)
}
