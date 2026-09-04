import { useState } from 'react'
import styles from "./App.module.css"
import * as about from '../../data/about.json'
import Header from './components/Header'

function getPreferredColorScheme() {
  if (!window.matchMedia) return 'light';

  if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
    return 'dark';
  }
  
  return 'light';
}

function App() {
  const [theme, setTheme] = useState<'light' | 'dark'>(getPreferredColorScheme())

  return <div className={styles[theme]}>
    <Header about={about}/>
  </div>
}

export default App
