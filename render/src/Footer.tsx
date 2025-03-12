import Contact from './Contact';

export default function Footer() {
    return (
        // So if we are on CV or Aticles we set the color to text-pink-600
        // Let say home pages is articles
        <footer className="flex py-2 px-20 items-center bg-emerald-200 bottom-0 flex-col mt-auto">
            <Contact />
        </footer>
        // Nav Bar
        // Liste des articles -> Loop sur une liste d'articles
        // Fetch depuis la db
    );
}
