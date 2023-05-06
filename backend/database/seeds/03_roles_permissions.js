/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> } 
 */
exports.seed = async function(knex) {
  // Deletes ALL existing entries
  await knex('public.roles_permissions').del()
  await knex('public.roles_permissions').insert([
    {id: 1, id_role: 1, id_permission: 1},
    {id: 2, id_role: 1, id_permission: 2},
    {id: 3, id_role: 2, id_permission: 1},
  ]);
};
