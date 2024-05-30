create table if not exists loans (
    id uuid primary key default gen_random_uuid(),
    user_id uuid not null,
    amount numeric not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    foreign key (user_id) references users (id)
);