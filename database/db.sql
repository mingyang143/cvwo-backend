create table discussion(
    ID int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id int NOT NULL,
    TITLE varchar(255) NOT NULL,
    CONTENT varchar(255) NOT NULL,
    LIKES int NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(ID)
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




