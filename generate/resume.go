package main

type Resume struct {
	About     ResumeAbout     `json:"about"`
	Skills    ResumeSkills    `json:"skills"`
	Education ResumeEducation `json:"education"`
}

type ResumeAbout struct {
	Name      ResumeName   `json:"name"`
	Positions []string     `json:"positions"`
	Links     []ResumeLink `json:"links"`
	Contact   []ResumeLink `json:"contact"`
}

type ResumeSkills map[string][]string

type ResumeEducation map[string]struct {
	Title       string     `json:"title"`
	Location    string     `json:"location"`
	Institution string     `json:"institution"`
	Date        ResumeDate `json:"date"`
	Notes       []string   `json:"notes"`
	Courses     []string   `json:"courses"`
}

type ResumeExperience map[string]struct {
	Title    string     `json:"title"`
	Location string     `json:"location"`
	Company  string     `json:"company"`
	Date     ResumeDate `json:"date"`
	Notes    []string   `json:"notes"`
	Skills   []string   `json:"skills"`
}

type ResumeProjects map[string]struct {
	Title  string     `json:"title"`
	Id     string     `json:"id"`
	Role   string     `json:"role"`
	Github string     `json:"github"`
	Date   ResumeDate `json:"date"`
	Notes  []string   `json:"notes"`
	Skills []string   `json:"skills"`
}

type ResumeName struct {
	First string `json:"First"`
	Last  string `json:"Last"`
}

type ResumeLink struct {
	Icon string  `json:"icon"`
	Text string  `json:"text"`
	Url  *string `json:"url"`
}

type ResumeDate struct {
	Start string `json:"start"`
	End   string `json:"end"`
}
