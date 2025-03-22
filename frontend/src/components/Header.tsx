import React from 'react';
import Navbar from '../components/sections/Navbar.tsx';

type Props = {}

const Header = (props: Props) => {
  return (
  <header className="bg-black text-2xl text-white">
    <h1>Weightlifting Library</h1>
    <Navbar />
  </header>
  )
}

export default Header
