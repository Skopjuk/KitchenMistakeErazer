CREATE TABLE IF NOT EXISTS recipes
(
    id serial primary key,
    recipe_name varchar,
    description varchar,
    recipe_version_id bigint,
    sourness int,
    saltiness int,
    acidity int,
    sweetness int,
    hot int,
    calories int,
    fat int,
    protein int,
    carbs int,
    created_at              timestamp not null default current_timestamp,
    updated_at              timestamp not null default current_timestamp
);

CREATE TABLE IF NOT EXISTS recipe_versions
(
    id serial primary key,
    recipe_name varchar,
    description varchar,
    recipe_id bigint,
    user_id bigint,
    sourness int,
    saltiness int,
    acidity int,
    sweetness int,
    hot int,
    calories int,
    fat int,
    protein int,
    carbs int,
    created_at              timestamp not null default current_timestamp,
    updated_at              timestamp not null default current_timestamp,
    constraint fk_user_id
        foreign key(user_id)
            references kitchen_users(id),
    constraint fk_recipe_id
        foreign key(id)
            references recipes(id)
);

