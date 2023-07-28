import Head from 'next/head'
import { Inter } from 'next/font/google'
import Button from '../componenets/button'
import { BUTTON_TYPES } from '@/common/data/button'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <div className='home'>
      <div className='btn-wrapper'>
        <Button type={`${BUTTON_TYPES.PRIMARY}`} btnText="Primary" />
        <Button type={`${BUTTON_TYPES.SECONDARY}`} btnText="Secondary" />
        <Button type={`${BUTTON_TYPES.TERTIARY}`} btnText="Tertiary" />
      </div>

    </div>
  )
}
