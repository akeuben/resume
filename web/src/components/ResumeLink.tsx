import type { Info } from "../types/About";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

export default function ResumeLink({icon, text, url}: Info) {
    return <a href={url}>
        <FontAwesomeIcon icon={["fas", icon]} />
        <span>{text}</span>
    </a>
}