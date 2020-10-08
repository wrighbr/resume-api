package models

type ContactInfo struct {
	Name     string
	Mobile   string
	Email    string
	Address  string
	Town     string
	Country  string
	Github   string
	LinkedIn string
	Website  string
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
