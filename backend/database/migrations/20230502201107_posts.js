/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.up = function(knex) {
    return knex.schema.withSchema('public').createTable("posts", table => {
        table.increments('id', { primaryKey: true })
        table.integer('id_author').notNullable().comment('Идентификатор автора поста')
        table.foreign('id_author').references('id').inTable('users')
        table.string("title", 100).notNullable().comment("Название поста")
        table.text("annotation").notNullable().comment("Аннотация поста")
        table.text("text").notNullable().comment("Основной текст поста")
        table.string('picture', 200).notNullable().comment('Картинка поста')
        table.bigInteger('views').notNullable().defaultTo(0).checkPositive()
            .comment('Просмотры поста')
        table.boolean('approved').notNullable().defaultTo(false)
            .comment('Флаг, одобрен ли пост')
        table.boolean('rejected').notNullable().defaultTo(false).comment('Флаг, отклонен ли пост')
        table.timestamp('time_start').nullable().comment('Время начала доступности')
        table.timestamp('time_end').nullable().comment('Время конца доступности')
        table.timestamp('created_at').notNullable().defaultTo(knex.fn.now())
            .comment('Время создания поста')
    })
};

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.down = function(knex) {
    return knex.schema.withSchema('public').dropTableIfExists('posts')
};
