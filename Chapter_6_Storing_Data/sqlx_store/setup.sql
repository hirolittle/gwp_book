drop table if exists posts cascade ;

create table posts (
    id      serial primary key,
    content text,
    author  varchar(255)
);
