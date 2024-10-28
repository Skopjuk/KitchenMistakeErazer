CREATE TABLE IF NOT EXISTS recipes
(
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT current_timestamp,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
            REFERENCES kitchen_users(id)

);

CREATE TABLE IF NOT EXISTS recipe_versions
(
    id SERIAL PRIMARY KEY,
    recipe_name VARCHAR(50) NOT NULL,
    description VARCHAR,
    recipe_version_number BIGINT NOT NULL,
    recipe_id BIGINT NOT NULL,
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
