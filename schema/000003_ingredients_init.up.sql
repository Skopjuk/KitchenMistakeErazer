CREATE TABLE IF NOT EXISTS measurements
(
    id SERIAL PRIMARY KEY,
    unit VARCHAR NOT NULL,
    CONSTRAINT unit_idx UNIQUE
        (unit)
);

CREATE TABLE IF NOT EXISTS ingredients
(
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    amount FLOAT,
    measurement_unit_id int,
    CONSTRAINT fk_measurement_unit_id
        FOREIGN KEY (measurement_unit_id)
            REFERENCES measurements(id)
);

CREATE TABLE IF NOT EXISTS recipe_ingredient
(
    id SERIAL PRIMARY KEY,
    recipe_id BIGINT,
    ingredient_id BIGINT,
    CONSTRAINT fk_ingredient_id
        FOREIGN KEY (ingredient_id)
            REFERENCES ingredients(id),
    CONSTRAINT fk_recipe_id
        FOREIGN KEY (recipe_id)
            REFERENCES recipes(id)
);