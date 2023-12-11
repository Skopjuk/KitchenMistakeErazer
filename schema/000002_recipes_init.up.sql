CREATE TABLE IF NOT EXISTS recipe_versions
(
    id serial primary key,
    recipe_name varchar,
    description varchar,
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
            references kitchen_users(id)
);

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
        updated_at              timestamp not null default current_timestamp,
        constraint fk_recipe_version_id
            foreign key(recipe_version_id)
                references recipe_versions(id)
);
