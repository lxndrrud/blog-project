/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.up = function(knex) {
    return knex.schema.withSchema('public').createTable("users", table => {
        table.increments('id', { primaryKey: true })
        table.string("login", 40).notNullable().comment('Логин пользователя')
        table.string("password", 200).notNullable().comment('Пароль пользователя')
        table.integer('id_role').notNullable().comment('Идентификатор роли')
        table.foreign('id_role').references('id').inTable('roles')
    })
};

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.down = function(knex) {
    return knex.schema.withSchema('public').dropTableIfExists('users')
  
};
