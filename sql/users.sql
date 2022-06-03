create table users (
	fullname varchar(250),
	email varchar(250) unique primary key,
	phone varchar(13) unique,
	password_hash varchar(2000),
	isVerified boolean,
	address varchar(2000)
);


CREATE TABLE Persons (
    Personid int NOT NULL AUTO_INCREMENT,
    LastName varchar(255) NOT NULL,
    FirstName varchar(255),
    Age int,
    PRIMARY KEY (Personid)
); 