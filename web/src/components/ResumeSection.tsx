import styles from "./ResumeSection.module.css"
import type { BulletItem, Item } from "../types/Section.d.ts"
import ResumeItem from "./ResumeItem.tsx"

export function ResumeSection({title, items}: {title: string, items: Item[]}) {
    return <section className={styles.section}>
        <h2>{title}</h2>
        {items.map((item, i) => <ResumeItem key={i} item={item}/>)}
    </section>
}

export function ResumeSectionBullet({title, items}: {title: string, items: BulletItem[]}) {
    return <section className={styles.section}>
        <h2>{title}</h2>
        <ul>
            {
                items.map(({title, values}) => <li key={title}><b>{title[0].toUpperCase() + title.substring(1)}: </b>{values.join(", ")}</li>)
            }
        </ul>
    </section>
}
