create table discussion(
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id int NOT NULL,
    title varchar(255) NOT NULL,
    content MEDIUMTEXT NOT NULL,
    likes int NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

create table comments(
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    comment MEDIUMTEXT NOT NULL,
    discussion_id int NOT NULL,
    FOREIGN KEY (discussion_id) REFERENCES discussion(id)
);

create table users(
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username varchar(255) NOT NULL,
    CONSTRAINT UN_USER_CONSTRAINT UNIQUE(username)
);

create table tags(
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    tag varchar(255) NOT NULL,
    CONSTRAINT UN_TAGS_CONSTRAINT UNIQUE(tag)
);

create table discussion_tags(
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    discussion_id int NOT NULL,
    tag_id int NOT NULL,
    FOREIGN KEY (discussion_id) REFERENCES discussion(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id),
    CONSTRAINT UN_DISCUSSIONTAG_CONSTRAINT UNIQUE(discussion_id,tag_id)

);




