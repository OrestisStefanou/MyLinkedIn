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

//Get profile photo url
func (prof *Professional) getPhotoURL() string {
	photoURL := mediaURL + prof.Photo
	return photoURL
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
