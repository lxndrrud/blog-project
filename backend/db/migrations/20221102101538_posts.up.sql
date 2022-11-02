CREATE TABLE IF NOT EXISTS "public"."posts" (
    id SERIAL PRIMARY KEY,
    id_author INT NOT NULL,
    title VARCHAR(100) NOT NULL,
    text TEXT NOT NULL,
    CONSTRAINT fk_author
        FOREIGN KEY (id_author)
            REFERENCES "public"."users"(id)
);

COMMENT ON COLUMN posts.id_author IS 'Идентификатор автора';
COMMENT ON COLUMN posts.title IS 'Заголовок поста';
COMMENT ON COLUMN posts.text IS 'Текст поста';
