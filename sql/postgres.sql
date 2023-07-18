CREATE TABLE users (   
    id serial primary key,
    name varchar(50) NOT NULL,
    nick varchar(50) NOT NULL unique,
    email varchar(50) NOT NULL unique,
    password varchar(20) NOT NULL unique,
    created_at timestamp default current_timestamp
);