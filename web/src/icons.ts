import { fas } from "@fortawesome/free-solid-svg-icons";
import { far } from "@fortawesome/free-regular-svg-icons";
import { fab } from "@fortawesome/free-brands-svg-icons";
import type { IconDefinition } from "@fortawesome/fontawesome-common-types";

export const icons = {
  ...fas,
  ...far,
  ...fab,
};

export const lookupIcon = (icon: string) => Object.values(icons).find(i => i.iconName === icon) as IconDefinition
