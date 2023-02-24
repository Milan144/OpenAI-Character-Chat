package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func main() {

	// Connect to phpmyadmin mysql database
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/open-character-chat")
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	} else {
		fmt.Println("Connected to database")
	}

	// Tables definitions

	/*
	 *   Game table
	 */
	type Game struct {
		id            int    `ID:"id"`
		title         string `Title:"title"`
		releaseDate   string `Release date:"releaseDate"`
		isMultiplayer bool   `Is multiplayer:"isMultiplayer"`
	}
	/*
	 *   User table
	 */
	type User struct {
		id       int    `ID:"id"`
		username string `Username:"username"`
	}
	/*
	 *   Message table
	 */
	type Message struct {
		id       int    `ID:"id"`
		content  string `Content:"content"`
		datetime string `Date:"date"`
	}
	/*
	 *   GameCharacter table
	 */
	type gameCharacter struct {
		id          int    `ID:"id"`
		name        string `Name:"name"`
		personality string `Personality:"personality"`
		game        int    `Game:"game"`
		image       string `Image:"image"`
	}
	/*
	 *   Conversations table
	 */
	type Conversation struct {
		id          int    `ID:"id"`
		userId      int    `User ID:"userId"`
		characterId int    `Character ID:"characterId"`
		isOpen      bool   `Is open:"isOpen"`
		lastMsgDate string `Last message date:"lastMsgDate"`
	}

	// GAMES ROUTES //

	// GET ALL games
	http.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		getImage("Gangplank", "League of Legends")
		log.Println("GET all games")
		results, err := db.Query("SELECT * FROM `Game`")
		if err != nil {
			panic(err.Error())
		}
		for results.Next() {
			var game Game
			err = results.Scan(&game.id, &game.title, &game.releaseDate, &game.isMultiplayer)
			if err != nil {
				panic(err.Error())
			}
			fmt.Fprintf(w, "ID: %d, Title: %s, Release date: %s, Is multiplayer: %t \n", game.id, game.title, game.releaseDate, game.isMultiplayer)
		}
		defer results.Close()
	})

	// GET ONE game
	http.HandleFunc("/game/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		log.Println("GET a game by id", id)
		results, err := db.Query("SELECT * FROM `Game` WHERE id = ?", id)
		if err != nil {
			panic(err.Error())
		}
		for results.Next() {
			var game Game
			err = results.Scan(&game.id, &game.title, &game.releaseDate, &game.isMultiplayer)
			if err != nil {
				panic(err.Error())
			}
			fmt.Fprintf(w, "ID: %d, Title: %s, Release date: %s, Is multiplayer: %t \n", game.id, game.title, game.releaseDate, game.isMultiplayer)
		}
		defer results.Close()
	})

	// POST add a new game
	http.HandleFunc("/game/add/", func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		releaseDate := r.URL.Query().Get("releaseDate")
		isMultiplayer := r.URL.Query().Get("isMultiplayer")
		log.Println("POST add a new game", title, releaseDate, isMultiplayer)
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
		id := r.URL.Query().Get("id")
		title := r.URL.Query().Get("title")
		releaseDate := r.URL.Query().Get("releaseDate")
		isMultiplayer := r.URL.Query().Get("isMultiplayer")
		log.Println("PUT update a game", id, title, releaseDate, isMultiplayer)
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
		id := r.URL.Query().Get("id")
		log.Println("DELETE a game", id)
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
		if err != nil {
			panic(err.Error())
		}
		for results.Next() {
			var character gameCharacter
			err = results.Scan(&character.id, &character.name, &character.personality, &character.game, &character.image)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			fmt.Fprintf(w, "ID: %d, Name: %s, Personality: %s, Game: %d \n", character.id, character.name, character.personality, character.game)
		}
		defer results.Close()
	})

	// GET ONE character
	http.HandleFunc("/character/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		log.Println("GET a character by id", id)
		results, err := db.Query("SELECT * FROM `gameCharacter` WHERE id = ?", id)
		if err != nil {
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
		name := r.URL.Query().Get("name")
		game := r.URL.Query().Get("game")

		gameName := db.QueryRow("SELECT title FROM Game WHERE id = ?", game)
		var gameTitle string
		err := gameName.Scan(&gameTitle)
		if err != nil {
			panic(err.Error())
		}

		personality := getInfos(name, gameTitle)
		//image := getImage(name, gameTitle)
		fmt.Println("Getting image for " + name + " from " + gameTitle)
		//image := getImage(name, gameTitle)
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
		id := r.URL.Query().Get("id")
		name := r.URL.Query().Get("name")
		personality := r.URL.Query().Get("personality")
		game := r.URL.Query().Get("game")
		log.Println("PUT update a character", id, name, personality, game)
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
		id := r.URL.Query().Get("id")
		log.Println("DELETE a character", id)
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
	// Get all conversations
	http.HandleFunc("/conversations", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET all conversations")
		results, err := db.Query("SELECT * FROM `conversations`")
		if err != nil {
			panic(err.Error())
		}
		for results.Next() {
			var conversation Conversation
			err = results.Scan(&conversation.id, &conversation.userId, &conversation.characterId, &conversation.isOpen, &conversation.lastMsgDate)
			if err != nil {
				panic(err.Error())
			}
			fmt.Fprintf(io.Writer(w), conversation.lastMsgDate)
		}

		defer results.Close()
	})

	// POST send a message to the chatbot and get the response
	http.HandleFunc("/openai/chatbot/", func(w http.ResponseWriter, r *http.Request) {

		characterId := r.URL.Query().Get("character")
		character := db.QueryRow("SELECT * FROM gameCharacter WHERE id = ?", characterId)
		var id int
		var name string
		var personality string
		var game string
		err = character.Scan(&id, &name, &personality, &game)
		if err != nil {
			panic(err.Error())
		}
		question := "In a roleplay and fun context, talk to me in english like u are  " + name + " from the game " + game + " : " + personality

		// Load .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		// Get api key from .env file
		apiKey := os.Getenv("API_KEY")

		// Create a new request to openAI API
		engineID := "text-davinci-003"
		url := fmt.Sprintf("https://api.openai.com/v1/engines/%s/completions", engineID)

		requestData := struct {
			Prompt    string `json:"prompt"`
			MaxTokens int    `json:"max_tokens"`
			MaxLength int    `json:"max_length"`
		}{Prompt: question, MaxTokens: 100, MaxLength: 500}

		requestBody, _ := json.Marshal(requestData)

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

		client := &http.Client{}
		resp, _ := client.Do(req)

		var responseData OpenAIResponse
		json.NewDecoder(resp.Body).Decode(&responseData)

		answer := responseData.Choices[0].Text
		fmt.Fprintf(w, answer)
	})

	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)

	defer db.Close()
}

// GET the infos about a character
func getInfos(name string, game string) string {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Get api key from .env file
	apiKey := os.Getenv("API_KEY")

	engineID := "text-ada-003"
	url := fmt.Sprintf("https://api.openai.com/v1/engines/%s/completions", engineID)

	question := "Resume me the story of " + name + " from " + game + "in 100 words"

	requestData := struct {
		Prompt string `json:"prompt"`
	}{Prompt: question}

	requestBody, _ := json.Marshal(requestData)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, _ := client.Do(req)

	var responseData OpenAIResponse

	json.NewDecoder(resp.Body).Decode(&responseData)

	answer := responseData.Choices[0].Text
	return answer
}

func getImage(name string, game string) string {
	println("Get image for " + name + " from " + game)
	url := "https://api.openai.com/v1/images/generations"

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Get api key from .env file
	apiKey := os.Getenv("API_KEY")

	// create JSON payload
	jsonData := []byte(`{
        "prompt": " ` + name + ` from ` + game + ` with a synthwave style",
        "n": 1,
        "size": "1024x1024"
    }`)

	// create request object with headers and payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// send HTTP request and get response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	// read response body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// print response body
	fmt.Println(string(body))
	return string(body)
}
