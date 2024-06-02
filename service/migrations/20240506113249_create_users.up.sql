create table if not exists users (
    id uuid primary key default gen_random_uuid(),
    name varchar(255) not null,
    email varchar(255) not null,
    password varchar(255) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);