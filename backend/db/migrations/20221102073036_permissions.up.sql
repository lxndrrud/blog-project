CREATE TABLE IF NOT EXISTS "public"."permissions" (
    id SERIAL PRIMARY KEY,
    title VARCHAR(60) NOT NULL,
    code VARCHAR(60) NOT NULL
);

COMMENT ON COLUMN permissions.title IS 'Название разрешения';
COMMENT ON COLUMN permissions.code IS 'Код разрешения';
