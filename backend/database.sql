CREATE TABLE Professionals (
	ProfessionalID int NOT NULL AUTO_INCREMENT,
	First_Name varchar(100) NOT NULL,
	Last_Name varchar(100) NOT NULL,
	Email varchar(255) NOT NULL,
	Password varchar(255) NOT NULL,
	Phone_Number varchar(100) NOT NULL,
	Photo varchar(255) NOT NULL,
	PRIMARY KEY(ProfessionalID)
);
