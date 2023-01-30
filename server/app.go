package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
    // Load the .env file and get the API_KEY
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    apiKey := os.Getenv("API_KEY")
    println(apiKey)

    // Connect to phpmyadmin mysql database
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/open-character-chat");

    // Check for errors
    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println("Connected to database");
    }

    fmt.Println("Server running on port 8080")

    // Tables definitions
    
    /*
    *   Game table
    */
    type Game struct {
      id   int    `ID:"id"`
      title string `Title:"title"`
      releaseDate string `Release date:"releaseDate"`
      isMultiplayer bool `Is multiplayer:"isMultiplayer"`
    }
    /*
    *   User table
    */
    type User struct {
      id   int    `ID:"id"`
      username string `Username:"username"`
    }
    /*
    *   Message table
    */
    type Message struct {
      id   int    `ID:"id"`
      content string `Content:"content"`
      datetime string `Date:"date"`
      isSentByHuman bool `Is sent by human:"isSentByHuman"`
    }
    /*
    *   GameCharacter table
    */
    type gameCharacter struct {
      id   int    `ID:"id"`
      name string `Name:"name"`
      personality string `Personality:"personality"`
      game int `Game:"game"`
    }
    /*
    *   Conversations table
    */
    type Conversation struct {
      id   int    `ID:"id"`
      userId int `User ID:"userId"`
      characterId int `Character ID:"characterId"`
      isOpen bool `Is open:"isOpen"`
      lastMsgDate string `Last message date:"lastMsgDate"`
    }

    // GAMES ROUTES //

    // GET ALL games 
    http.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
      log.Println("GET all games")
      results, err := db.Query("SELECT * FROM `Game`")
      if err !=nil {
          panic(err.Error())
      }
      for results.Next() {
        var game Game
        err = results.Scan(&game.id, &game.title, &game.releaseDate, &game.isMultiplayer)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
        fmt.Fprintf(io.Writer(w), game.title, game.releaseDate, game.isMultiplayer)
      } 

      defer results.Close()
    })

    // GET a game by id 
    http.HandleFunc("/game/", func(w http.ResponseWriter, r *http.Request) {
      id := r.URL.Query().Get("id")
      log.Println("GET a game by id", id)
      // Get the id from the url
      results, err := db.Query("SELECT * FROM `Game` WHERE id = ?", id)
      if err !=nil {
          panic(err.Error())
      }
      for results.Next() {
        var game Game
        err = results.Scan(&game.id, &game.title, &game.releaseDate, &game.isMultiplayer)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
        fmt.Fprintf(io.Writer(w), game.title, game.releaseDate, game.isMultiplayer)
      }
      defer results.Close()
    })

    // POST add a new game
    http.HandleFunc("/game/add/", func(w http.ResponseWriter, r *http.Request) {
      // Getting parameters
      title := r.URL.Query().Get("title")
      releaseDate := r.URL.Query().Get("releaseDate")
      isMultiplayer := r.URL.Query().Get("isMultiplayer")
      
      log.Println("POST add a new game", title, releaseDate, isMultiplayer)

      // Example url
      // http://localhost:8080/game/add/?title=Call of Duty Modern Warfare 2&releaseDate=01-10-2022&isMultiplayer=1

      // Inserting data 
      stmt, err := db.Prepare("INSERT INTO Game (title, releaseDate, isMultiplayer) VALUES (?,?,?)")
      if err != nil {
          panic(err.Error())
      }
      log.Println("INSERT INTO Game", title, releaseDate, isMultiplayer)
      _, err = stmt.Exec(title, releaseDate, isMultiplayer)
      if err != nil {
         panic(err.Error())
      }

      fmt.Println("Game added successfully")
    })

    // CHARACTERS ROUTES //
    
    // GET all characters
    http.HandleFunc("/characters", func(w http.ResponseWriter, r *http.Request) {
      log.Println("GET all characters")
      results, err := db.Query("SELECT * FROM `gameCharacter`")
      if err !=nil {
          panic(err.Error())
      }
      for results.Next() {
        var character gameCharacter
        err = results.Scan(&character.id, &character.personality, &character.game)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
        fmt.Fprintf(io.Writer(w), character.personality, character.game)
      } 

      defer results.Close()
    })
    // GET a character by id 
    http.HandleFunc("/characters/:id", func(w http.ResponseWriter, r *http.Request) {
      id := r.URL.Query().Get("id")
      log.Println("GET a character by id", id)
      // Get the id from the url
      results, err := db.Query("SELECT * FROM `gameCharacter` WHERE id = ?", id)
      if err !=nil {
          panic(err.Error())
      }
      for results.Next() {
        var character gameCharacter
        err = results.Scan(&character.id, &character.personality, &character.game)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
        fmt.Fprintf(io.Writer(w), character.personality, character.game)
      }
      defer results.Close()
    })
    // POST add a new character 
    http.HandleFunc("/characters/add", func(w http.ResponseWriter, r *http.Request) {
      // Getting parameters
      // TODO : Change request
      name := r.URL.Query().Get("name")
      personality := r.URL.Query().Get("personality")
      game := r.URL.Query().Get("game")


      log.Println("POST add a new game", title, releaseDate, isMultiplayer)

      // Example url
      // http://localhost:8080/game/add/?title=Call of Duty Modern Warfare 2&releaseDate=01-10-2022&isMultiplayer=1

      // Inserting data 
      stmt, err := db.Prepare("INSERT INTO gameCharacter (personality, game) VALUES (?,?)")
      if err != nil {
          panic(err.Error())
      }
      log.Println("INSERT INTO Game", title, releaseDate, isMultiplayer)
      _, err = stmt.Exec(title, releaseDate, isMultiplayer)
      if err != nil {
         panic(err.Error())
      }

      fmt.Println("Character added successfully")
    })
    
    // CONVERSATIONS ROUTES //
    http.HandleFunc("/conversations", func(w http.ResponseWriter, r *http.Request) {
      log.Println("GET all conversations")
      results, err := db.Query("SELECT * FROM `conversations`")
        if err !=nil {
            panic(err.Error())
        }
        for results.Next() {
          var conversation Conversation
          err = results.Scan(&conversation.id, &conversation.userId, &conversation.characterId, &conversation.isOpen, &conversation.lastMsgDate)
          if err != nil {
              panic(err.Error()) // proper error handling instead of panic in your app
          }
          fmt.Fprintf(io.Writer(w), conversation.lastMsgDate)
        } 

        defer results.Close()
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

    defer db.Close()
}
