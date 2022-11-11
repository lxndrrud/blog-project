CREATE TABLE IF NOT EXISTS "public"."posts" (
    id SERIAL PRIMARY KEY,
    id_author INT NOT NULL,
    title VARCHAR(100) NOT NULL,
    annotation TEXT NOT NULL,
    text TEXT NOT NULL,
    picture VARCHAR(200) NOT NULL,
    views BIGINT NOT NULL DEFAULT 0 CHECK (views >= 0),
    approved BOOLEAN NOT NULL DEFAULT FALSE,
    rejected BOOLEAN NOT NULL DEFAULT FALSE,
    time_start TIMESTAMP,
    time_end TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_author
        FOREIGN KEY (id_author)
            REFERENCES "public"."users"(id)
);

COMMENT ON COLUMN posts.id_author IS 'Идентификатор автора';
COMMENT ON COLUMN posts.title IS 'Заголовок поста';
COMMENT ON COLUMN posts.annotation IS 'Аннотация поста';
COMMENT ON COLUMN posts.text IS 'Текст поста';
COMMENT ON COLUMN posts.picture IS 'Картинка поста';
COMMENT ON COLUMN posts.views IS 'Количество просмотров';
COMMENT ON COLUMN posts.approved IS 'Флаг, проверен ли пост';
COMMENT ON COLUMN posts.time_start IS 'Время начала доступности';
COMMENT ON COLUMN posts.time_end IS 'Время конца доступности';
COMMENT ON COLUMN posts.created_at IS 'Время и дата создания поста';


