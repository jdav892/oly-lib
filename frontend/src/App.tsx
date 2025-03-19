import { useState } from 'react';
import './App.css';
import Header from "./components/Header.tsx";
import Footer from "./components/Footer.tsx";
import Body from "./components/Body.tsx";


function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <Header />
      <Body />
      <Footer />
   </>
  )
}

export default App
