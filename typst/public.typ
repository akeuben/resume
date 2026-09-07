#import "./resume.typ": *

#let (about, education, experience, skills, projects) = json("../data/resume.json")

#resume(about, education, experience, skills, projects)
