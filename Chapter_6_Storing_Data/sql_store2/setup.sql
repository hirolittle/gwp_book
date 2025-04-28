DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS comments;

create table posts (
    id      serial primary key,
    content text,
    author  varchar(255)
);

create table comments (
    id      serial primary key,
    content text,
    author  varchar(255),
    post_id integer references posts(id)
);