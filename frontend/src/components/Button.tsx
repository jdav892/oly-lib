import React from 'react'

type Props = {
  title: string
}
{/*Needs to be reusable*/}
function ({ title }: Props) {
  return (
  <button>{title}</button>
  )
}

export default 
