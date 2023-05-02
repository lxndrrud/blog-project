/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.up = function(knex) {
    return knex.schema.withSchema('public').createTable("roles_permissions", table => {
        table.increments('id', { primaryKey: true })
        table.integer('id_role').notNullable().comment('Идентификатор роли')
        table.foreign('id_role').references('id').inTable('roles')
        table.integer('id_permission').notNullable().comment('Идентификатор разрешения')
        table.foreign('id_permission').references('id').inTable('permissions')
    })
};

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.down = function(knex) {
    return knex.schema.withSchema('public').dropTableIfExists('roles_permissions')
};
