/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.up = function(knex) {
    return knex.schema.withSchema('public').createTable("permissions", table => {
        table.increments('id', { primaryKey: true })
        table.string("title", 60).notNullable().comment('Название разрешения')
        table.string("code", 60).notNullable().comment('Код разрешения')
    })
};

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.down = function(knex) {
    return knex.schema.withSchema('public').dropTableIfExists('permissions')
};
