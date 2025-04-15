import React from 'react';
import Navbar from '../components/sections/Navbar.tsx';

type Props = {}

const Header = (props: Props) => {
  return (
  <header>
    <h1 className="flex items-center justify-center">Weightlifting Library</h1>
      <Navbar />
  </header>
  )
}

export default Header
