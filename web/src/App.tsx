import { useState } from 'react'
import styles from "./App.module.css"
import about from '../../data/about.json'
import skills from '../../data/skills.json'
import education from '../../data/education.json'
import experience from '../../data/experience.json'
import projects from '../../data/projects.json'
import ResumeHeader from './components/ResumeHeader'
import Header from './components/Header'
import { ResumeSectionBullet, ResumeSection } from './components/ResumeSection'
import ResumeLink from './components/ResumeLink'
import type { Item } from "./types/Section.d.ts"

function getPreferredColorScheme() {
  if (!window.matchMedia) return 'light';

  if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
    return 'dark';
  }
  
  return 'light';
}

function formatDate(date) {
    return `${date.start} - ${date.end || "present"}`
}

const transformers = {
    education: (data) => data.map(data => ({
        title: data.title,
        subtitle: data.institution,
        location: data.location,
        date: formatDate(data.date),
        content: <ul>
            {data.notes.map(note => <li key={note}>{note}</li>)}
            <li><b>Notable Courses: </b>{data.courses.join(", ")}</li>
        </ul>
    })),
    experience: (data) => data.map(data => ({
        title: data.title,
        subtitle: data.company,
        location: data.location,
        date: formatDate(data.date),
        content: <ul>
            {data.notes.map(note => <li key={note}>{note}</li>)}
            <li><b>Skills: </b>{data.skills.join(", ")}</li>
        </ul>
    })),
    projects: (data) => data.map(data => ({
        title: data.page ? <a href={`https://akeuben.ca/project/` + data.page}>data.title</a> : data.title,
        subtitle: data.role,
        location: data.github ? <ResumeLink icon={"github"} text={data.github} url={`https://github/com/${data.github}`} /> : "",
        date: data.date,
        content: <ul>
            {data.notes.map(note => <li key={note}>{note}</li>)}
            {data.skills && <li><b>Skills: </b>{data.skills.join(", ")}</li>}
        </ul>
    }))
}


function App() {
  const [theme, setTheme] = useState<'light' | 'dark'>(getPreferredColorScheme())

  return <div className={[styles.body, styles[theme]].join(" ")}>
    <Header theme={theme} setTheme={setTheme} />
    <main className={styles.resume}>
      <ResumeHeader about={about}/>
      <ResumeSectionBullet title="Skills" items={skills} />
      <ResumeSection title="Education" items={transformers.education(education)} />
      <ResumeSection title="Experience" items={transformers.experience(experience)} />
      <ResumeSection title="Projects" items={transformers.projects(projects)} />
    </main>
  </div>
}

export default App
