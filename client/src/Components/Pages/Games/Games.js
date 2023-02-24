import React, { useState, useEffect } from 'react';
import Navbar1 from "../../Subcomponents/Navbar";
import axios from 'axios';
import './Games.css';

export default function Games(props) {

    const [games, setGames] = useState([]);
    const getAllGames = () => {
        axios.get(`http://localhost:8000/games`)
            .then((response) => {
                const allGames = response.data;
                const gamesArray = allGames.split('\n').map(game => {
                    const [id, title, releaseDate] = game.split(', ');

                    const gameId = id ? id.split(': ')[1] : null;
                    const gameTitle = title ? title.split(': ')[1] : null;
                    const gameReleaseDate = releaseDate ? releaseDate.split(': ')[1] : null;

                    return {
                        id: gameId,
                        title: gameTitle,
                        releaseDate: gameReleaseDate,
                    };
                });
                console.log(gamesArray)
                setGames(gamesArray);
            })
            .catch((error) => console.error(`Error: ${error}`));
    }

    useEffect(() => {
        console.log('calling getAllgames')
        getAllGames();
    }, []);

    return (
          <div>
              <Navbar1 />
              <h1 className="pageTitle">List of games</h1>

              <div className="container">
                  <main className="grid">
                      {games.map((game, index) => {
                          return (
                              <div key={index}>
                                  <article>
                                          <div className="text">
                                              <h3>{game.title}</h3>
                                              <p>{game.releaseDate}</p>
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
