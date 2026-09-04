import type { About } from "../types/About";
import styles from "./Header.module.css"

export default function Header({about}: {about: About}) {
  return <header className={styles.header}>
    <h1>{about.name.first} {about.name.last}</h1>
    <p>{
        about.positions.join(" · ")
    }</p>
  </header>
}