package data

type DisplayableListItem interface {
	Display() string
}

type Resume struct {
	About      ResumeAbout        `json:"about"`
	Skills     []ResumeSkill      `json:"skills"`
	Education  []ResumeEducation  `json:"education"`
	Experience []ResumeExperience `json:"experience"`
	Projects   []ResumeProjects   `json:"projects"`
}

type ResumeAbout struct {
	Name      ResumeName   `json:"name"`
	Positions []string     `json:"positions"`
	Links     []ResumeLink `json:"links"`
	Contact   []ResumeLink `json:"contact"`
}

type ResumeSkill struct {
	Title  string   `json:"title"`
	Values []string `json:"values"`
}

func (self ResumeSkill) Display() string {
	return self.Title
}

type ResumeEducation struct {
	Title       string     `json:"title"`
	Location    string     `json:"location"`
	Institution string     `json:"institution"`
	Date        ResumeDate `json:"date"`
	Notes       []string   `json:"notes"`
	Courses     []string   `json:"courses"`
}

func (self ResumeEducation) Display() string {
	return self.Title
}

type ResumeExperience struct {
	Title    string     `json:"title"`
	Location string     `json:"location"`
	Company  string     `json:"company"`
	Date     ResumeDate `json:"date"`
	Notes    []string   `json:"notes"`
	Skills   []string   `json:"skills"`
}

func (self ResumeExperience) Display() string {
	return self.Title
}

type ResumeProjects struct {
	Title  string   `json:"title"`
	Id     string   `json:"id"`
	Role   string   `json:"role"`
	Github string   `json:"github"`
	Date   int      `json:"date"`
	Notes  []string `json:"notes"`
	Skills []string `json:"skills"`
}

func (self ResumeProjects) Display() string {
	return self.Title
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
