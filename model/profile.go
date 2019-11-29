package model

type Profile struct {
	Name string
	Gender string
	Age string
	Height string
	Weight string
	Income string
	Marriage string
	WorkPlace string
	Work string
	Education string
}

func (p Profile) String() string {
	return "Name: " + p.Name + "\n Gender: " + p.Gender + "\n Age: " + p.Age + "\n Height: " +
		p.Height + "\n Weight: " + p.Weight + "\n Income: " + p.Income + "\n Marriage: " + p.Marriage +
		"\n WorkPlace: " + p.WorkPlace + "\n Work: " + p.Work + "\n Education: " + p.Education
}