import styles from "./ResumeSection.module.css"
import type { Item } from "../types/Section.d.ts"
import ResumeItem from "./ResumeItem.tsx"

export function ResumeSection({title, items}: {title: string, items: Item[]}) {
    return <section className={styles.section}>
        <h2>{title}</h2>
        {items.map(item => <ResumeItem key={item.title} item={item}/>)}
    </section>
}

export function ResumeSectionBullet({title, items}: {title: string, items: Record<string, string[]>}) {
    return <section className={styles.section}>
        <h2>{title}</h2>
        <ul>
            {
                Object.keys(items).map(title => <li key={title}><b>{title[0].toUpperCase() + title.substring(1)}: </b>{items[title].join(", ")}</li>)
            }
        </ul>
    </section>
}
