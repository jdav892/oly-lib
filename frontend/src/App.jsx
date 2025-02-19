import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import Header from "./components/Header.jsx"
import Footer from "./components/Footer.jsx"
import Display from "./components/Display.jsx"

function App()  
  const [count, setCount] = useState(0)

  return (
    <main>
      <Header />
      <Display />
      <Footer />
    </main>
  )
}

export default App
