CREATE TABLE recipes
(
    id SERIAL PRIMARY KEY,
    user_id BIGINT,
    created_at  TIMESTAMP NOT NULL DEFAULT current_timestamp,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
            REFERENCES kitchen_users(id)

);

CREATE TABLE recipe_versions
(
    id SERIAL PRIMARY KEY,
    recipe_name VARCHAR(50),
    description VARCHAR,
    recipe_version_number BIGINT,
    recipe_id BIGINT,
    sourness INT,
    saltiness INT,
    acidity INT,
    sweetness INT,
    hot INT,
    calories INT,
    fat INT,
    protein INT,
    carbs INT,
    created_at              TIMESTAMP NOT NULL DEFAULT current_timestamp,
    CONSTRAINT fk_recipe_id
        FOREIGN KEY (recipe_id)
            REFERENCES recipes(id)
);

