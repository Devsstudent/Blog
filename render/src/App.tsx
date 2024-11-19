import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

import Home from './Home.tsx';

import {Routes, Route} from 'react-router-dom';

// On essaie de limite un max les inputs utilisateur

// On peux stocker tt les post dans le context, au chargement de la page principal

// 

function App() {
  // Fetch les post depuis la db 
  // Les afficher sur la page main
  // Sur le click on load la page du post en question
  const [count, setCount] = useState(0);

  debugger;

  return (
    <>
      <Routes>
        <Route path="/test" element={<Home />} />
      </Routes>
      <div>
          BRUH
      </div>
    </>
  )
}

export default App
