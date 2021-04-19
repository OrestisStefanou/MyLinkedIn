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
