import React from 'react'
import Weightlifting from './sections/Weightlifting';
import Powerlifting from './sections/Powerlifting';
import Fitness from './sections/Fitness';
import Plyometrics from './sections/Plyometrics';


type Props = {}

const Body = (props: Props) => {
  return (
    <section className="grid grid-cols-5">
      <><Weightlifting /><Powerlifting /><Fitness /><Plyometrics /></>
    </section>
  )
}

export default Body
