import styles from "./Header.module.css"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { lookupIcon } from '../icons'

export default function Header({theme, setTheme}: {theme: 'light' | 'dark', setTheme: (arg0: 'light' | 'dark') => void}) {
    return <header className={styles.header}>
        <h1>&lt;Avery<br />&nbsp;Keuben&gt;</h1>
        <a href="https://akeuben.ca">Back to Site</a>
        <Print />
        <ToggleTheme theme={theme} setTheme={setTheme} />
    </header>
}

function Print() {
    return <button onClick={() => {
        window.location.replace("https://github.com/akeuben/resume/releases/latest/download/resume.pdf");
    }}>
        <FontAwesomeIcon icon={lookupIcon("print")} />
    </button>
}

function ToggleTheme({theme, setTheme}: {theme: 'light' | 'dark', setTheme: (arg0: 'light' | 'dark') => void}) {
    return <button onClick={() => {
        setTheme(theme === 'light' ? 'dark' : 'light')
    }}>
        <FontAwesomeIcon icon={lookupIcon(theme === 'light' ? "moon" : "sun")} />
    </button>
}
