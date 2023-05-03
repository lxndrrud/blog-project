import Redis from "ioredis";

export const RedisConnection = new Redis({
    host: 'cacher-blog',
    port: 7379,
    password: 'redis_password'
})
