from flask import Flask, jsonify
import requests
from dotenv import load_dotenv
load_dotenv()
app = Flask(__name__)
import os

# TODO: Connect to the database 
# TODO: Connect to openAI API and send request / get response

# Get the api key from the .env file
api_key = os.getenv("API_KEY")

# ROUTES

# Home
@app.route('/')
def index():
    return 'Welcome to the Flask API'

# GET List of all characters
@app.route('/characters')
def get_characters():
    # Get the data from the database

    # Convert the data to JSON
    data = response.json()
    # Return the data
    return jsonify(data)

# POST Create a new character with the name as parameter of the function
@app.route('/characters/<name>')
def create_character(name):
    # Create a new character in the database
    
    # Return the new character
    return jsonify(response.json())
   
# GET a character by id
@app.route('/characters/<id>')
def get_character(id):
    # Get the character from the database

    # Return the character
    return jsonify(response.json())

# GET the infos of the character
@app.route('/characters/<id>/infos')
def get_character_infos(id):
    # Get the character infos from the database

    # Return the character infos
    return jsonify(response.json())

# Talk to the character using the AI 
@app.route('/characters/<id>/talk/<message>')
def talk_to_character(id, message):
    # Talk to the character using the AI

    # Return the answer of the AI
    return jsonify(response.json())

# Run the app

if __name__ == '__main__':
       app.run(debug=True)
