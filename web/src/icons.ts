import { fas } from "@fortawesome/free-solid-svg-icons";
import { far } from "@fortawesome/free-regular-svg-icons";
import { fab } from "@fortawesome/free-brands-svg-icons";

export const icons = {
  ...fas,
  ...far,
  ...fab,
};

export const lookupIcon = (icon) => Object.values(icons).find(i => i.iconName === icon)
