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

CREATE TABLE Article_Likes (
       id int NOT NULL AUTO_INCREMENT,
       ProfessionalID int NOT NULL,
       ArticleID int NOT NULL,
       PRIMARY KEY(id),
       FOREIGN KEY(ProfessionalID) REFERENCES Professionals(ProfessionalID),
       FOREIGN KEY(ArticleID) REFERENCES Articles(id)
);

CREATE TABLE Article_Comments (
       id int NOT NULL AUTO_INCREMENT,
       ProfessionalID int NOT NULL,
       ArticleID int NOT NULL,
       Comment varchar(255),
       Created datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
       PRIMARY KEY(id),
       FOREIGN KEY(ProfessionalID) REFERENCES Professionals(ProfessionalID),
       FOREIGN KEY(ArticleID) REFERENCES Articles(id)
);

CREATE TABLE Notifications (
       id int NOT NULL AUTO_INCREMENT,
       ProfessionalID int NOT NULL,
       Msg varchar(255),
       Seen boolean,
       Created datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
       PRIMARY KEY(id),
       FOREIGN KEY(ProfessionalID) REFERENCES Professionals(ProfessionalID)
);

/*Professional with ID1 sends a request to professional with ID2*/
CREATE TABLE Friendships (
       id int NOT NULL AUTO_INCREMENT,
       ProfessionalID1 int NOT NULL,
       ProfessionalID2 int NOT NULL,
       Status ENUM('pending', 'friends'),
       PRIMARY KEY(id),
       FOREIGN KEY(ProfessionalID1) REFERENCES Professionals(ProfessionalID),
       FOREIGN KEY(ProfessionalID2) REFERENCES Professionals(ProfessionalID)
);

CREATE TABLE Messages (
       id int NOT NULL AUTO_INCREMENT,
       Sender int NOT NULL,
       Receiver int NOT NULL,
       Msg TEXT NOT NULL,
       Seen boolean,
       Created datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
       PRIMARY KEY(id),
       FOREIGN KEY(Sender) REFERENCES Professionals(ProfessionalID),
       FOREIGN KEY(Receiver) REFERENCES Professionals(ProfessionalID)
);


SELECT * FROM Articles WHERE (UploaderID IN (SELECT ProfessionalID2 FROM Friendships WHERE ProfessionalID1 = 16) OR  UploaderID IN (SELECT ProfessionalID1 FROM Friendships WHERE ProfessionalID2=16)) 
OR (id IN(SELECT al.ArticleID FROM Article_Likes al,Friendships f WHERE al.ProfessionalID = f.ProfessionalID2 AND f.ProfessionalID1=16) OR id IN(SELECT al.ArticleID FROM Article_Likes al,Friendships f WHERE al.ProfessionalID = f.ProfessionalID1 AND f.ProfessionalID2=16)) ORDER BY Created DESC;
