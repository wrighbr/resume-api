package models

type ContactInfo struct {
	Name     string `json:"name,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Email    string `json:"email,omitempty"`
	Address  string `json:"address,omitempty"`
	Town     string `json:"town,omitempty"`
	Country  string `json:"country,omitempty"`
	Github   string `json:"github,omitempty"`
	LinkedIn string `json:"linkedin,omitempty"`
	Website  string `json:"website,omitempty"`
}

type Experience struct {
	StartDate      string
	EndDate        string
	Company        string
	Role           string
	Description    string
	Duration       string
	Location       string
	EmploymentType string
}

type Education struct {
	School    string
	Sourse    string
	StartDate string
	EndDate   string
}

type Language struct {
	Language    string
	Proficiency string
}

type TechSkills struct {
}
