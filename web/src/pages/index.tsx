import Head from 'next/head'
import { Inter } from 'next/font/google'
import Button from '../componenets/button'
import NavBar from '../componenets/navbar'
import { BUTTON_TYPES } from '@/common/data/button'
import background from '../../public/JER09793.jpg'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <div
      style={{
        backgroundImage: `url(${background.src})`,
        backgroundPosition: 'center',
        backgroundSize: 'cover',
        backgroundRepeat: 'no-repeat',
        width: '100vw',
        height: '100vh',
      }}>

      <NavBar />
      <div className='home'>
        <div>Greenpointe</div>

      </div>
      <div className='btnWrapper'>
        <Button type={`${BUTTON_TYPES.PRIMARY}`} btnText="View Recipe" />
        <Button type={`${BUTTON_TYPES.SECONDARY}`} btnText="Search Recipes" />
      </div>
    </div>
  )
}
