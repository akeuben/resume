import type { Item } from "../types/Section.d.ts"
import styles from "./ResumeItem.module.css"

export default function ResumeItem({item}: {item: Item}) {
    return <div className={styles.item}>
        <div className={styles.row}>
            <b>{item.title}</b>
            <b>{item.location}</b>
        </div>
        <div className={styles.row}>
            <p>{item.subtitle}</p>
            <p>{item.date}</p>
        </div>
        {item.content}
    </div>
}
