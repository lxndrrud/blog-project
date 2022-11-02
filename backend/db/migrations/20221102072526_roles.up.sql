CREATE TABLE IF NOT EXISTS "public"."roles" (
    id SERIAL PRIMARY KEY,
    title VARCHAR(30) NOT NULL
);

COMMENT ON COLUMN roles.title IS 'Название роли';

