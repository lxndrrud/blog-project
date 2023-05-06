/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> } 
 */
exports.seed = async function(knex) {
  // Deletes ALL existing entries
  await knex('public.roles').del()
  await knex('public.roles').insert([
    {id: 1, title: 'Администратор'},
    {id: 2, title: 'Блогер'},
  ]);
};
