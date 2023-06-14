import Head from 'next/head'
import { Inter } from 'next/font/google'
import Button from '../componenets/common/Button'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <>
      <Head>
        <title>Jeffy's Bar</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Button variant='primary'>Hey</Button>
      <Button variant='secondary'>Hey</Button>
    </>
  )
}
