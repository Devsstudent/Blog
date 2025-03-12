import AddArticle from './AddArticle.tsx';
import Home from './Home.tsx';
import NavBar from './Navbar.tsx';
import { Routes, Route } from 'react-router-dom';
import RenderArticle from './RenderArticle.tsx';
import Articles from './Articles.tsx';
import Cv from './Cv.tsx';
import Contact from './Contact.tsx';
import Footer from './Footer.tsx';

//import {Routes, Route} from 'react-router-dom';

// On essaie de limite un max les inputs utilisateur

// On peux stocker tt les post dans le context, au chargement de la page principal

//

export default function App() {
    // Fetch les post depuis la db
    // Les afficher sur la page main
    // Sur le click on load la page du post en question

    return (
        <div className="flex flex-col min-h-screen">
            <NavBar />
            <div className="flex-1 bg-[#e6fff9]">
                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="/add" element={<AddArticle />} />
                    <Route path="/article/:title" element={<RenderArticle />} />
                    <Route path="/contact" element={<Contact />} />
                    <Route path="/cv" element={<Cv />} />
                    <Route path="/articles" element={<Articles />} />
                </Routes>
            </div>
            <Footer />
        </div>
    );
}
