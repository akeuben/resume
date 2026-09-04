export type About = {
    name: {
        first: string,
        last: string,
    },
    positions: string[],
    links: Info[],
    contact: Info[],
}

export type Info = {
    icon: string,
    text: string,
    url?: string,
}
