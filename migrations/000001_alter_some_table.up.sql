CREATE TABLE if not exists students(
    id serial primary key,
    first_name varchar(30) not null,
    last_name varchar(30) not null,
    username varchar(30) not null,
    email varchar(255) not null,
    phone_number varchar(255),
    created_at TIMESTAMP WITH TIME ZONE default current_timestamp
);