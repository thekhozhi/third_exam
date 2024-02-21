CREATE TABLE if not exists books (
    id uuid primary key,
    name varchar(100) not null, 
    author_name varchar(60), 
    page_number integer,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at integer default 0
);

INSERT INTO books (id,name, author_name, page_number) values ('a0c5d953-529b-4340-a907-753ae828dcae', 'The richest man', 'Khojiakbar',
130);


