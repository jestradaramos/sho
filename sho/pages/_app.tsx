import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { Jacques_Francois } from '@next/font/google'


const jf = Jacques_Francois({weight: '400', subsets: ['latin']})

export default function App({ Component, pageProps }: AppProps) {
  return (
    <main className={jf.className}>
      <Component {...pageProps} />
    </main>
    
  )
}
