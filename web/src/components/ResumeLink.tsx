import type { Info } from "../types/About";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

import { fas } from "@fortawesome/free-solid-svg-icons";
import { far } from "@fortawesome/free-regular-svg-icons";
import { fab } from "@fortawesome/free-brands-svg-icons";

const icons = {
  ...fas,
  ...far,
  ...fab,
};

export default function ResumeLink({icon, text, url}: Info) {
    const svg = Object.values(icons).find(i => i.iconName === icon);
    console.log(text);
    return <a href={url}>
        <FontAwesomeIcon icon={svg} />
        <span>{text}</span>
    </a>
}
