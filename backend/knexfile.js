// Update with your config settings.

const path = require('path');

/**
 * @type { Object.<string, import("knex").Knex.Config> }
 */
module.exports = {
  main: {
    client: 'pg',
    searchPath: [ 'public' ],
    connection: {
        host : 'postgres-blog',
        port : 5432,
        user : 'user',
        password : 'password',
        database : 'blog'
    },
    migrations: {
        tableName: 'migrations',
        directory: path.join(__dirname, 'database/migrations')
    },
    seeds: {
        directory: path.join(__dirname, 'database/seeds')
      },
    pool: {
      min: 2,
      max: 10
    },
  }
};
