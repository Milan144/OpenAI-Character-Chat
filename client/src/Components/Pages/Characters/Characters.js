import React, { useState, useEffect } from 'react';
import Navbar1 from "../../Subcomponents/Navbar";
import axios from 'axios';
import './Characters.css';

export default function Characters(props) {

    const [characters, setCharacters] = useState([]);
    const getAllCharacters = () => {
        axios.get(`http://localhost:8000/characters`)
            .then((response) => {
                const allCharacters = response.data;
                const charactersArray = allCharacters.split('\n').map(game => {
                    const [id, name, personality] = game.split(', ');

                    const characterId = id ? id.split(': ')[1] : null;
                    const characterName = name ? name.split(': ')[1] : null;
                    const characterPersonality = personality ? personality.split(': ')[1] : null;

                    return {
                        id: characterId,
                        name: characterName,
                        personality: characterPersonality,
                    };
                });
                console.log(charactersArray)
                setCharacters(charactersArray);
            })
            .catch((error) => console.error(`Error: ${error}`));
    }

    useEffect(() => {
        console.log('calling getAllCharacters')
        getAllCharacters();
    }, []);

    return (
        <div>
            <Navbar1 />
            <h1 className="pageTitle">List of characters</h1>

            <div className="container">
                <main className="grid">
                    {characters.map((game, index) => {
                        return (
                            <div key={index}>
                                <article>
                                    <div className="text">
                                        <h3>{game.name}</h3>
                                        <p>{game.personality}</p>
                                    </div>
                                </article>
                            </div>
                        )
                    })}
                </main>
            </div>






        </div>
    );
}
