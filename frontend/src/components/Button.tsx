import React from 'react'

type Props = {
  title: string
}
{/*Needs to be reusable*/}
function ({ title }: Props) {
  return (
    <button className="rounded-md bg-black text-white">
      {title}
    </button>
  )
}

export default 
