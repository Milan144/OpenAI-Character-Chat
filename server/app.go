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
        // Format result to display it on page
        fmt.Fprintf(w, "ID: %d, Title: %s, Release date: %s, Is multiplayer: %t \n", game.id, game.title, game.releaseDate, game.isMultiplayer)

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
        fmt.Fprintf(w, "ID: %d, Title: %s, Release date: %s, Is multiplayer: %t \n", game.id, game.title, game.releaseDate, game.isMultiplayer)


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

    // PUT update a game
    http.HandleFunc("/game/update/", func(w http.ResponseWriter, r *http.Request) {
      // Getting parameters
      id := r.URL.Query().Get("id")
      title := r.URL.Query().Get("title")
      releaseDate := r.URL.Query().Get("releaseDate")
      isMultiplayer := r.URL.Query().Get("isMultiplayer")

      log.Println("PUT update a game", id, title, releaseDate, isMultiplayer)

      // Example url
      // http://localhost:8080/game/update/?id=1&title=Call of Duty Modern Warfare 2&releaseDate=01-10-2022&isMultiplayer=1

      // Updating data
      stmt, err := db.Prepare("UPDATE Game SET title = ?, releaseDate = ?, isMultiplayer = ? WHERE id = ?")
      if err != nil {
          panic(err.Error())
      }
      log.Println("UPDATE Game SET title = ?, releaseDate = ?, isMultiplayer = ? WHERE id = ?", title, releaseDate, isMultiplayer, id)
      _, err = stmt.Exec(title, releaseDate, isMultiplayer, id)
      if err != nil {
          panic(err.Error())
      }

      fmt.Println("Game updated successfully")
    })

    // DELETE a game
    http.HandleFunc("/game/delete/", func(w http.ResponseWriter, r *http.Request) {
      // Getting parameters
      id := r.URL.Query().Get("id")

      log.Println("DELETE a game", id)

      // Example url
      // http://localhost:8080/game/delete/?id=1

      // Deleting data
      stmt, err := db.Prepare("DELETE FROM Game WHERE id = ?")
      if err != nil {
          panic(err.Error())
      }
      log.Println("DELETE FROM Game WHERE id = ?", id)
      _, err = stmt.Exec(id)
      if err != nil {
          panic(err.Error())
      }

      fmt.Println("Game deleted successfully")
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
        err = results.Scan(&character.id, &character.name, &character.personality, &character.game)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
        fmt.Fprintf(w, "ID: %d, Name: %s, Personality: %s, Game: %d \n", character.id, character.name, character.personality, character.game)
      } 

      defer results.Close()
    })
    // GET a character by id 
    http.HandleFunc("/character/", func(w http.ResponseWriter, r *http.Request) {
      id := r.URL.Query().Get("id")
      log.Println("GET a character by id", id)
      // Get the id from the url
      results, err := db.Query("SELECT * FROM `gameCharacter` WHERE id = ?", id)
      if err !=nil {
          panic(err.Error())
      }
      for results.Next() {
        var character gameCharacter
        err = results.Scan(&character.id, &character.name, &character.personality, &character.game)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
         fmt.Fprintf(w, "ID: %d, Name: %s, Personality: %s, Game: %d \n", character.id, character.name, character.personality, character.game)
      }
      defer results.Close()
    })

    // POST add a new character 
    http.HandleFunc("/character/add/", func(w http.ResponseWriter, r *http.Request) {
      // Getting parameters
      // TODO : Change request
      name := r.URL.Query().Get("name")
      personality := r.URL.Query().Get("personality")
      game := r.URL.Query().Get("game")


      // TODO: get personality parameter with a request to openAI api

      log.Println("POST add a new game", name, personality, game)

      // Example url
      // http://localhost:8080/game/add/?title=Call of Duty Modern Warfare 2&releaseDate=01-10-2022&isMultiplayer=1

      // Inserting data 
      stmt, err := db.Prepare("INSERT INTO gameCharacter (name, personality, game) VALUES (?,?,?)")
      if err != nil {
          panic(err.Error())
      }
      _, err = stmt.Exec(name, personality, game)
      if err != nil {
         panic(err.Error())
      }

      fmt.Println("Character added successfully")
    })

    // PUT update a character
    http.HandleFunc("/character/update/", func(w http.ResponseWriter, r *http.Request) {
      // Getting parameters
      id := r.URL.Query().Get("id")
      name := r.URL.Query().Get("name")
      personality := r.URL.Query().Get("personality")
      game := r.URL.Query().Get("game")

      log.Println("PUT update a character", id, name, personality, game)

      // Example url
      // http://localhost:8080/game/update/?id=1&title=Call of Duty Modern Warfare 2&releaseDate=01-10-2022&isMultiplayer=1

      // Updating data
      stmt, err := db.Prepare("UPDATE gameCharacter SET name = ?, personality = ?, game = ? WHERE id = ?")
      if err != nil {
          panic(err.Error())
      }
      _, err = stmt.Exec(name, personality, game, id)
      if err != nil {
          panic(err.Error())
      }

      fmt.Println("Character updated successfully")
    })

    // Delete a character
    http.HandleFunc("/character/delete/", func(w http.ResponseWriter, r *http.Request) {
      // Getting parameters
      id := r.URL.Query().Get("id")
      
      log.Println("DELETE a character", id)
      
      // Deleting data
      stmt, err := db.Prepare("DELETE FROM gameCharacter WHERE id = ?")
      if err != nil {
          panic(err.Error())
      }
      _, err = stmt.Exec(id)
      if err != nil {
          panic(err.Error())
      }

      fmt.Println("Character deleted successfully")
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
    http.HandleFunc("/openai/character", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "GET the infos about a character with his game and name")
    })
    
    // POST send a message to the chatbot and get the response
    http.HandleFunc("/openai/chatbot", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "POST send a message to the chatbot and get the response")
    })

	  http.ListenAndServe(":8080", nil)

    defer db.Close()
}
