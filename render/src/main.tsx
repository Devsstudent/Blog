import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { BrowserRouter } from 'react-router-dom'
import BlogContextProvider from './BlogContext.tsx'
import App from './App.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <BlogContextProvider >
          <App />
      </BlogContextProvider>
    </BrowserRouter>
  </StrictMode>,
)
