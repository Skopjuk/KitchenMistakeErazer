CREATE TABLE IF NOT EXISTS kitchen_users
(
    id         serial primary key,
    first_name  varchar(255) not null,
    last_name  varchar(255) not null,
    email      varchar(255) not null,
    password   varchar(255) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    constraint users_email_idx unique
        (email)
);