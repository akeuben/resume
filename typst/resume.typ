#import "@preview/fontawesome:0.6.2": *
#import "@preview/decasify:0.11.4": *

#let gap = 0.25em

#let sep = " " + sym.dot.c + " "

#let resume_link(..args) = {
  let icon = args.at(0)
  let text = args.at(1)
  let target = args.at(2)
  if target != none {
    return fa-icon(icon) + " " + link(target, text)
  } else {
    return fa-icon(icon) + " " + text
  }
}

#let resume_header(content) = box(stroke: (bottom: 1pt), inset: (bottom: 2pt), width: 1fr, text(weight: "bold", content))
#let resume-item(item) = {
  let title = item.at("title", default: "")
  let subtitle = item.at("subtitle", default: "")
  let location = item.at("location", default: "")
  let date = item.at("date", default: "")
  let content = item.at("content", default: "")
  return box(inset: (bottom: gap))[
    #text(weight: "bold", [#title #h(1fr) #location])\
    #if subtitle != "" {
      text(weight: "medium", [#subtitle #h(1fr) #date])
    }
    #content
  ]
}

#let resume-date-parse(date) = {
  let end = date.at("end", default: "present")

  return date.start + " - " + end
}

#let resume-education-parse(item) = {
  return (
    title: item.title,
    subtitle: item.institution,
    location: item.location,
    date: resume-date-parse(item.date),
    content: [
      #list(..item.notes, [*Notable Courses:* #item.courses.join(", ")])
    ]
  )
}

#let resume-experience-parse(item) = {
  return (
    title: item.title,
    subtitle: item.company,
    location: item.location,
    date: resume-date-parse(item.date),
    content: [
      #list(..item.notes, [*Skills:* #item.skills.join(", ")])
    ]
  )
}

#let resume-project-parse(item) = {
  return (
    title: if item.at("page", default: none) != none {link("https://akeuben.ca/project/" + item.page, item.title)} else {item.title},
    subtitle: item.role,
    date: item.date,
    location: if item.at("github", default: none) != none {resume_link("github", item.github, "https://github.com/" + item.github)} else {[]},
    content: if item.at("skills", default: none) != none {list(..item.notes, [*Skills:* #item.skills.join(", ")])} else {list(..item.notes)}
  )
}


#let resume(about, education, experience, skills, projects) = {
  set page(margin: 0.33in)
  set par(spacing: gap)
  set document(title: [Resume - #about.name.first #about.name.last])
  
  [
    #align(center)[
      #text(size: 3em, weight: "bold", font: "Roboto", about.name.first + " " + about.name.last)
      
      #text(size: 1em, weight: "bold", about.at("positions").join(sep))
      
      #about.at("links").map(link => resume_link(link.icon, link.text, link.url)).join(sep)
      
      #about.at("contact").map(link => resume_link(link.icon, link.text, link.at("url", default: none))).join(sep)
    ]

    #resume_header("Skills")\
    #box(inset: (bottom: gap), list(..skills.keys().map(category => [*#titlecase(category):* #skills.at(category).join(", ")])))
  
    #resume_header("Education")
    #education.map(item => resume-item(resume-education-parse(item))).join()

    #resume_header("Experience")
    #experience.map(item => resume-item(resume-experience-parse(item))).join()

    #resume_header("Projects")
    
    #let project_url = about.links.find(i => i.icon == "link").url + "/projects";
    More projects on my website, #link(project_url)
    #projects.map(item => resume-item(resume-project-parse(item))).join()

    #resume_header("References")
    
    Available Upon Request
  ]
}