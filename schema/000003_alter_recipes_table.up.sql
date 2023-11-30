ALTER TABLE recipes ADD column user_id BIGINT;

ALTER TABLE recipes ADD constraint fk_user_id
    FOREIGN KEY (user_id)
        REFERENCES kitchen_users(id);
