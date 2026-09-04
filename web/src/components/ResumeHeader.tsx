import { Fragment } from 'react';
import type { About } from "../types/About";
import styles from "./ResumeHeader.module.css"
import ResumeLink from "./ResumeLink";

export default function ResumeHeader({about}: {about: About}) {
  return <header className={styles.header}>
    <h1>{about.name.first} {about.name.last}</h1>
    <p>{
        about.positions.join(" · ")
    }</p>
    <div className={styles.linkgroup}>
        {about.links.map((link, i) => <Fragment key={i}>
            <ResumeLink icon={link.icon} text={link.text} url={link.url}/>
            {i < about.links.length - 1 && <span> · </span>}
        </Fragment>)}
    </div>
    <div className={styles.linkgroup}>
        {about.contact.map((link, i) => <Fragment key={i}>
            <ResumeLink icon={link.icon} text={link.text} url={link.url}/>
            {i < about.contact.length - 1 && <span> · </span>}
        </Fragment>)}
    </div>
  </header>
}
