-- create database if not exists devdb;
-- use devdb;

drop table if exists users;

create table users(
    id serial not null unique,
    name varchar(255) not null,
    email varchar(255) not null unique,
    password varchar(255) not null,
    created_at timestamp default current_timestamp,
    primary key(id)
);

-- insert into users(name, email, password)
-- values
--     ()
