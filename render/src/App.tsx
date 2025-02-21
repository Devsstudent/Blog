import AddArticle from './AddArticle.tsx';
import Home from './Home.tsx';
import NavBar from './Navbar.tsx';
import { Route, Routes } from 'react-router-dom';

//import {Routes, Route} from 'react-router-dom';

// On essaie de limite un max les inputs utilisateur

// On peux stocker tt les post dans le context, au chargement de la page principal

// 

export default function App() {
  // Fetch les post depuis la db 
  // Les afficher sur la page main
  // Sur le click on load la page du post en question

  return (
    <>
      <NavBar />
      <Routes>
        <Route path='/' element={<Home />} />
        <Route path='/add' element={< AddArticle />} />
      </Routes>
      <div className="text-purple-500">
        OK
      </div>
    </>)
}