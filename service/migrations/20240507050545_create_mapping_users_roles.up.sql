create table if not exists mapping_users_roles (
    user_id uuid not null,
    role_id uuid not null,
    primary key (user_id, role_id),
    foreign key (user_id) references users (id),
    foreign key (role_id) references roles (id)
);