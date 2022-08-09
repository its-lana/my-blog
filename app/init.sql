CREATE DATABASE myblog;

use myblog;

create table posts (
  id serial not null unique,
  title varchar(64),
  category varchar(64),
  content text,
  primary key(id)
);

insert into posts(title, category, content)
values
    ('Review Smartphone Samsung A03s', 'Smartphone','The obligatory Hello World Post ...'),
    ('Review Laptop Lenovo Legion Slim S7', 'laptop', 'Yet another blog post about something exciting');


CREATE DATABASE myblog_test;

use myblog_test;

create table posts (
  id serial not null unique,
  title varchar(64),
  category varchar(64),
  content text,
  primary key(id)
);

DROP TABLE posts;

