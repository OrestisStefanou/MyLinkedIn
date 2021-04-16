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

//Create a save method for Professional
func (prof *Professional) save() error {
	err := dbclient.createProfessional(prof)
	return err
}
