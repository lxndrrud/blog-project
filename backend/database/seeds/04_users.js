/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> } 
 */
exports.seed = async function(knex) {
  // Deletes ALL existing entries
  await knex('public.users').del()
  await knex('public.users').insert([
    { 
      login: 'admin', 
      password: '$2b$10$vUcvAnO0XtwO2x.K9X3xX.3pEH17R9aJdCmAfZitLmT3DWPWxijgG',
      id_role: 1
    },
    {
      login: 'bloger', 
      password: '$2b$10$vUcvAnO0XtwO2x.K9X3xX.3pEH17R9aJdCmAfZitLmT3DWPWxijgG',
      id_role: 2
    },
  ]);
};
