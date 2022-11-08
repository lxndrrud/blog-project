CREATE TABLE IF NOT EXISTS "public"."users" (
    id SERIAL PRIMARY KEY,
    id_role INT NOT NULL,
    login VARCHAR(40) NOT NULL,
    password VARCHAR(200) NOT NULL,
    CONSTRAINT fk_role
        FOREIGN KEY (id_role)
            REFERENCES "public"."roles"(id)
);

COMMENT ON COLUMN users.login IS 'Логин пользователя';
COMMENT ON COLUMN users.password IS 'Пароль пользователя';
COMMENT ON COLUMN users.id_role IS 'Идентификатор роли';
