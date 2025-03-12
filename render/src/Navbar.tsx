import { useNavigate } from 'react-router-dom';
import Logo from '../public/orsonlogo.svg'
export default function NavBar() {


  let navigate = useNavigate();
  const routeChange = (path: string) => {
    navigate(path);
  }

  return (

    // So if we are on CV or Aticles we set the color to text-pink-600
    // Let say home pages is articles
    <div className="sticky flex justify-between py-2 grid-cols-4 px-20 items-center bg-emerald-200 top-0">
      <img src={Logo} alt="Logo" className='w-24 h-16 hover:cursor-pointer' onClick={() => routeChange('/')} />
      <span className='text-xl font-semibold text-pink-400 hover:text-pink-600 cursor-pointer' onClick={() =>
        routeChange('/cv')}>CV</span>
      <span className='text-xl font-semibold text-pink-600 cursor-pointer' onClick={() =>
        routeChange('/articles')
      }>Articles</span>
      <span className='text-xl font-semibold text-pink-400 hover:text-pink-600 cursor-pointer' onClick={() =>
        routeChange('/contact')
      }>Contact</span>
    </div >
    // Nav Bar
    // Liste des articles -> Loop sur une liste d'articles
    // Fetch depuis la db
  )
}
