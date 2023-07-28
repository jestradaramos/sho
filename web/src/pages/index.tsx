import Head from 'next/head'
import { Inter } from 'next/font/google'
import Button from '../componenets/button'
import NavBar from '../componenets/navbar'
import { BUTTON_TYPES } from '@/common/data/button'
import background from '../../public/JER09793.jpg'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <>

      <div className='base'
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

          <div className='btnWrapper'>
            <Button type={`${BUTTON_TYPES.PRIMARY}`} btnText="Primary" />
            <Button type={`${BUTTON_TYPES.SECONDARY}`} btnText="Secondary" />
            <Button type={`${BUTTON_TYPES.TERTIARY}`} btnText="Tertiary" />
          </div>

        </div>
      </div>
    </>






  )
}
