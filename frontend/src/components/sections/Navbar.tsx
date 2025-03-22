import React from 'react'

type Props = {}

const Navbar = (props: Props) => {
  return (
    <nav className="text-white">
      <h2 className="text-xl">Modalities</h2>
      <button
        title="Competition"
        className="rounded-md"
      />
      <button 
        title="Strength"
        className="rounded-md"
      />
      <button
        title="Plyometrics"
        className="rounded-md"
      />
      <button
        title="Accessory"
        className="rounded-md"
      />
    </nav>
  )
}

export default Navbar
