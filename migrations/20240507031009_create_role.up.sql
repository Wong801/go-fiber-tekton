create table if not exists roles (
    id uuid primary key default gen_random_uuid(),
    name varchar(255) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

insert into roles (name) values 
    ('loan:create_loan'), 
    ('loan:read_loan'), 
    ('loan:update_loan'), 
    ('loan:delete_loan'), 
    ('user:create_user'), 
    ('user:read_user'), 
    ('user:update_user'),
    ('role:create_role'),
    ('role:read_role'),
    ('role:update_role'),
    ('role:delete_role');