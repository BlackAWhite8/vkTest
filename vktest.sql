\c vktest;
create table actor(id serial primary key, name varchar(50), sex varchar(10), birthday date);
create table movie(id serial primary key, title varchar(150) check (length(title) >= 1), description text check(length(description)<=1000), realise_date date, score smallint check (score >= 0 and score <= 10));
create table web_user(id serial primary key, login varchar(30), pass varchar(50), is_admin boolean default FALSE);
insert into web_user(login,pass,is_admin) values ('administrator', 'admin', TRUE);
create table movie_actors(row_id serial primary key, title varchar(150) check (length(title) >= 1), actor_name varchar(50));
insert into movie_actors(title, actor_name) values ('blow', 'Johny Depp');
insert into movie_actors(title, actor_name) values ('cocain', 'Johny Depp');