/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> } 
 */
exports.seed = async function(knex) {
  // Deletes ALL existing entries
  await knex('public.posts').del()
  await knex('public.posts').insert([
    {
      title: 'test title', 
      annotation: 'test annotation', 
      text: 'test text', 
      id_author: 1, 
      picture: '/diretide.jpeg', 
      views: 0, 
      approved: 'TRUE', 
      time_start: null, 
      time_end: '2024-01-01T10:00:00'
    },
    {
      title: 'test title 2', 
      annotation: 'test annotation 2', 
      text: 'test text 2', 
      id_author: 1, 
      picture: '/diretide.jpeg', 
      views: 0, 
      approved: true, 
      time_start: null, 
      time_end: '2024-01-01T10:00:00'
    },
    {
      title: 'test title 3', 
      annotation: 'test annotation 3', 
      text: 'test text 3', 
      id_author: 1, 
      picture: '/diretide.jpeg', 
      views: 0, 
      approved: true, 
      time_start: null, 
      time_end: '2024-01-01T10:00:00'
    },
  ]);
};
