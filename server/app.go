import (
  "encoding/jsn"
  "net/http"
  "github.com/joho/godotenv"
)

err := godotenv.Load()
if err != nil {
    log.Fatal("Error loading .env file")
}

apiKey := os.Getenv("API_KEY")


func askQuestion(question string) ([]byte, error) {
  endpoint := "https://api.openai.com/v1/engines/davinci-codex/completions"
  payload := map[string]string{"prompt": question,
"max_tokens": "100"}
  payloadBytes, err := json.Marshal(payload)

  req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(payloadBytes))
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer "+apiKey)
  
  client := &http.Client{}
  res, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()

  responseBytes, _ := ioutil.ReadAll(res.Body)
  return responseBytes, nil
}

func main() {
    question := "What is the capital of France?"
    responseBytes, err := askQuestion(question)
    if err != nil {
        fmt.Println(err)
        return
    }
    var response map[string]interface{}
    json.Unmarshal(responseBytes, &response)
    fmt.Println(response)
}
