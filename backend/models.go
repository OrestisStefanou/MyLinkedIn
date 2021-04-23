package main

//Professional json struct
type Professional struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phoneNumber"`
	Photo       string `json:"photo"`
}

//Save method for Professional
func (prof *Professional) save() error {
	err := dbclient.createProfessional(prof)
	return err
}

//Update method for Professional
func (prof *Professional) update() error {
	err := dbclient.updateProfessional(prof)
	return err
}

//Get profile photo url
func (prof *Professional) setPhotoURL() {
	photoURL := mediaURL + prof.Photo
	prof.Photo = photoURL
}

//Add education info to a professional
func (prof *Professional) addEducation(educationInfo Education) error {
	educationInfo.ProfessionalID = prof.ID
	err := dbclient.createEducation(&educationInfo)
	return err
}

//Get the education of a professional
func (prof *Professional) getEducation() ([]Education, error) {
	education, err := dbclient.getProfessionalEducation(prof.ID)
	return education, err
}

//Remove education of a professional
func (prof *Professional) removeEducation(educationInfo Education) error {
	educationInfo.ProfessionalID = prof.ID
	err := dbclient.deleteProfessionalEducation(educationInfo)
	return err
}

//Add experience info to a professional
func (prof *Professional) addExperience(experienceInfo Experience) error {
	experienceInfo.ProfessionalID = prof.ID
	err := dbclient.createExperience(&experienceInfo)
	return err
}

//Get the experience of a Professional
func (prof *Professional) getExperience() ([]Experience, error) {
	experience, err := dbclient.getProfessionalExperience(prof.ID)
	return experience, err
}

//Remove experience of a Professional
func (prof *Professional) removeExperience(experienceInfo Experience) error {
	experienceInfo.ProfessionalID = prof.ID
	err := dbclient.deleteProfessionalExperience(experienceInfo)
	return err
}

//Education json struct
type Education struct {
	ID             int    `json:"id"`
	ProfessionalID int    `json:"professionalId"`
	DegreeName     string `json:"degreeName" binding:"required"`
	SchoolName     string `json:"schoolName" binding:"required"`
	StartDate      string `json:"startDate" binding:"required"`
	FinishDate     string `json:"finishDate"`
}

//Save method for Education
func (education *Education) save() error {
	err := dbclient.createEducation(education)
	return err
}

//Method to set id of Education object
func (education *Education) setID() error {
	id, err := dbclient.getEducationID()
	education.ID = id + 1
	return err
}

//Experience json struct
type Experience struct {
	ID             int    `json:"id"`
	ProfessionalID int    `json:"professionalId"`
	EmployerName   string `json:"employerName" binding:"required"`
	JobTitle       string `json:"jobTitle" binding:"required"`
	StartDate      string `json:"startDate" binding:"required"`
	FinishDate     string `json:"finishDate"`
}

//Save method for experience
func (experience *Experience) save() error {
	err := dbclient.createExperience(experience)
	return err
}

//Method to set id of Experience object
func (experience *Experience) setID() error {
	id, err := dbclient.getExperienceID()
	experience.ID = id + 1
	return err
}
