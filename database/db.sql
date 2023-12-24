create table discussion(
    ID int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    USER varchar(255),
    TITLE varchar(255) NOT NULL,
    CONTENT varchar(255)
);

/*
    ID int NOT NULL,
    LastName varchar(255) NOT NULL,
    FirstName varchar(255),
    Age int,
    PRIMARY KEY (ID,LastName)
*/

create table comments(
    ID int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    comment varchar(255) NOT NULL,
    discussion_id int NOT NULL,
     FOREIGN KEY (discussion_id) REFERENCES discussion(ID)
);

create table users(
    ID int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    USERNAME varchar(255) NOT NULL
);
