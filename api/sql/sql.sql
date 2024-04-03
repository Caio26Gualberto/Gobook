CREATE DATABASE IF NOT EXISTS golang;
USE golang;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(20) not null,
    creationTime timestamp default current_timestamp()
) ENGINE=INNODB;