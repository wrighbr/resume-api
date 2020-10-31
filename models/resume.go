package models

type Resume struct {
	ContactInfo ContactInfo  `json:"contact_info,omitempty"`
	Experience  []Experience `json:"experience,omitempty"`
	Education   []Education  `json:"education,omitempty"`
	Language    []Language   `json:"language,omitempty"`
	TechSkills  []TechSkills `json:"tech_skills,omitempty"`
}

type ContactInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Email    string `json:"email,omitempty"`
	Address  string `json:"address,omitempty"`
	Town     string `json:"town,omitempty"`
	Country  string `json:"country,omitempty"`
	Github   string `json:"github,omitempty"`
	LinkedIn string `json:"linkedin,omitempty"`
	Website  string `json:"website,omitempty"`
	PostCode int    `json:"postcode,omitempty"`
}

type Experience struct {
	ID             int    `json:"id"`
	StartDate      string `json:"start_date,omitempty"`
	EndDate        string `json:"end_data,omitempty"`
	Company        string `json:"company,omitempty"`
	Role           string `json:"role,omitempty"`
	Description    string `json:"description,omitempty"`
	Duration       string `json:"duration,omitempty"`
	Location       string `json:"location,omitempty"`
	EmploymentType string `json:"employment_type,omitempty"`
}

type Education struct {
	ID          int    `json:"id"`
	Institution string `json:"institution,omitempty"`
	Course      string `json:"course,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
}

type Language struct {
	ID          int    `json:"id"`
	Language    string `json:"language,omitempty"`
	Proficiency string `json:"proficiency,omitempty"`
}

type TechSkills struct {
	ID    int      `json:"id"`
	Skill string   `json:"skill,omitempty"`
	Tags  []string `json:"tags,omitempty"`
	Stars int      `json:"Stars,omitempty"`
	// Technology []string `json:"technology,omitempty"`
}

type Auth struct {
	Username string
	Password string
}
