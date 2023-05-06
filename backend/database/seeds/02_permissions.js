/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> } 
 */
exports.seed = async function(knex) {
  // Deletes ALL existing entries
  await knex('public.permissions').del()
  await knex('public.permissions').insert([
    {id: 1, title: 'Создавать посты', code: 'СОЗД_ПОСТЫ'},
    {id: 2, title: 'Модерировать посты', code: 'МОДЕР_ПОСТЫ'},
  ]);
};
