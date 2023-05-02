import knex from 'knex'

export const DatabaseConnection = knex({
    client: 'pg',
    searchPath: [ 'public' ],
    connection: {
        host : 'postgres-blog',
        port : 5432,
        user : 'user',
        password : 'password',
        database : 'blog',
    }
});
  