import React from 'react'

type Props = {}

const Navbar = (props: Props) => {
  return (
    <nav>
      <h2 className="text-xl">Modalities</h2>
      <div className="grid grid-cols-4">
          <div>
            <button
              className="border-solid rounded-md"
                >Competition</button>
          </div>
        <div>
            <button 
              className="rounded-md"
                >Strength</button>
          </div>
        <div>
          <button
            className="rounded-md"
              >Plyometrics</button>
        </div>
        <div>
          <button
            className="rounded-md"
              >Accessories</button>
        </div>
      </div>
    </nav>
  )
}

export default Navbar
