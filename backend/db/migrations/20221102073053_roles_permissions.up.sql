CREATE TABLE IF NOT EXISTS "public"."roles_permissions" (
    id SERIAL PRIMARY KEY,
    id_role INT NOT NULL,
    id_permission INT NOT NULL,
    CONSTRAINT fk_role
        FOREIGN KEY (id_role)
            REFERENCES "public"."roles"(id),
    CONSTRAINT fk_permission
        FOREIGN KEY (id_permission)
            REFERENCES "public"."permissions"(id)
);

COMMENT ON COLUMN roles_permissions.id_role IS 'Идентификатор роли';
COMMENT ON COLUMN roles_permissions.id_permission IS 'Идентификатор разрешения';

