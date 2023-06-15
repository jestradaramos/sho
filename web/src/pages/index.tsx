import Head from 'next/head'
import { Inter } from 'next/font/google'
import Button from '../componenets/common/Button'
import Image from 'next/image'
import Box from '@mui/material/Box'
import Stack from '@mui/material/Stack'
import NavBar from '@/componenets/common/NavBar'
import Typography from '@mui/material/Typography'

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
        fill
        style={{ objectFit: 'cover', position: 'absolute' }}
      />


      <NavBar></NavBar>

      <div className='MainAlignButtons'>
        <Stack justifyContent={'center'} direction='row' spacing={{ xs: 4, sm: 8 }}>
          <Button variant='primary'>View Recipe</Button>
          <Button variant='secondary'>Search Others</Button>
        </Stack>
      </div>


      <Typography variant="h4" component="div"  color='#F1DAC4'>
        GreenPointe
      </Typography>





    </>
  )
}
