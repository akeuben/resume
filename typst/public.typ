#import "./resume.typ": *

#let about = json("../data/about.json")
#let education = json("../data/education.json")
#let experience = json("../data/experience.json")
#let skills = json("../data/skills.json")
#let projects = json("../data/projects.json")

#resume(about, education, experience, skills, projects)
