from flask import Flask, jsonify
import requests
from dotenv import load_dotenv
load_dotenv()
app = Flask(__name__)
import os

# Get the api key from the .env file
api_key = os.getenv("API_KEY")

# Home
@app.route('/')
def index():
    return 'Welcome to the Flask API'

# ask something to openAI api
@app.route('/api')
def api():
    headers = {
        'Content-Type': 'application/json',
        'Authorization': f'Bearer {api_key}'
    }
    endpoint = 'https://api.openai.com/v1/engines/davinci-codex/completions'
    payload = {
        'prompt': 'What is the name of the french president ?',
        'max_tokens': 100
    }
    response = requests.post(endpoint, json=payload, headers=headers)
    # Get the answer from the response
    answer = response.json()['choices'][0]['text']
    return answer


if __name__ == '__main__':
       app.run(debug=True)
