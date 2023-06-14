import Head from 'next/head'
import { Inter } from 'next/font/google'
import Button from '../componenets/common/Button'
import Image from 'next/image'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <>
      <Head>
        <title>Jeffy's Bar</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Image 
        src='/JER09793.jpg'
        alt='Greenpointe'
        quality={100}
        style={{objectFit: "cover"}}
        fill/>
      

      <Button variant='primary'>Hey</Button>
      <Button variant='secondary'>Hey</Button>
      
      
    </>
  )
}
