from flask import Flask, jsonify
import requests
from dotenv import load_dotenv
load_dotenv()
app = Flask(__name__)
import os
api_key = os.getenv("API_KEY")

@app.route('/')
def index():
    return 'Welcome to the Flask API'

@app.route('/api')
def api():
    headers = {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer' + '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'
    }
    endpoint = 'https://api.openai.com/v1/engines/davinci-codex/completions'
    payload = {
        'prompt': 'What is the capital of France?',
        'temperature': 0.5,
        'max_tokens': 100
    }
    response = requests.post(endpoint, json=payload, headers=headers)
    return jsonify(response.json())


if __name__ == '__main__':
       app.run(debug=True)
