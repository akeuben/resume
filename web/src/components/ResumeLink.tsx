import type { Info } from "../types/About";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { lookupIcon } from "../icons"

export default function ResumeLink({icon, text, url}: Info) {
    console.log(text);
    return <a href={url}>
        <FontAwesomeIcon icon={lookupIcon(icon)} />
        <span>{text}</span>
    </a>
}
