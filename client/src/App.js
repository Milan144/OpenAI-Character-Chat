import './App.css';
import {Route, Routes} from "react-router-dom";

//Pages
import Home from './Components/Pages/Home/Home.js';
import Characters from './Components/Pages/Characters/Characters.js';
import Games from './Components/Pages/Games/Games.js';

function App() {
  return (
      <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/characters" element={<Characters />} />
          <Route path="/games" element={<Games />} />
      </Routes>
  );
}

export default App;
