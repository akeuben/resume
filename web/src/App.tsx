import { useState } from 'react'
import styles from "./App.module.css"
import * as about from '../../data/about.json'
import Header from './components/Header'

import { library } from '@fortawesome/fontawesome-svg-core'

/* import all the icons in Free Solid, Free Regular, and Brands styles */
import { fas } from '@fortawesome/free-solid-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { fab } from '@fortawesome/free-brands-svg-icons'

library.add(fas, far, fab)

function getPreferredColorScheme() {
  if (!window.matchMedia) return 'light';

  if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
    return 'dark';
  }
  
  return 'light';
}

function App() {
  const [theme, setTheme] = useState<'light' | 'dark'>(getPreferredColorScheme())

  return <div className={[styles.body, styles[theme]].join(" ")}>
    <main className={styles.resume}>
      <Header about={about}/>
    </main>
  </div>
}

export default App
