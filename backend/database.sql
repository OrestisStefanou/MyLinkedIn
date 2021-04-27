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

CREATE TABLE Education (
       id int NOT NULL AUTO_INCREMENT,
       ProfessionalID int NOT NULL,
       Degree_Name varchar(255) NOT NULL,
       School_Name varchar(255) NOT NULL,
       Start_Date varchar(100) NOT NULL,
       Finish_Date varchar(100),
       PRIMARY KEY(id),
       FOREIGN KEY(ProfessionalID) REFERENCES Professionals(ProfessionalID)
);

CREATE TABLE Experience (
       id int NOT NULL AUTO_INCREMENT,
       ProfessionalID int NOT NULL,
       Employer_Name varchar(255) NOT NULL,
       Job_Title varchar(255) NOT NULL,
       Start_Date varchar(100) NOT NULL,
       Finish_Date varchar(100),
       PRIMARY KEY(id),
       FOREIGN KEY(ProfessionalID) REFERENCES Professionals(ProfessionalID)
);

CREATE TABLE Skills (
       id int NOT NULL AUTO_INCREMENT,
       ProfessionalID int NOT NULL,
       Name varchar(255) NOT NULL,
       PRIMARY KEY(id),
       FOREIGN KEY(ProfessionalID) REFERENCES Professionals(ProfessionalID)
);

CREATE TABLE Articles (
       id int NOT NULL AUTO_INCREMENT,
       UploaderID int NOT NULL,
       Title varchar(255) NOT NULL,
       Content TEXT  NOT NULL,
       Attached_File varchar(255),
       Created datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
       PRIMARY KEY(id),
       FOREIGN KEY(UploaderID) REFERENCES Professionals(ProfessionalID)
);
