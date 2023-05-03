import { Redis } from "ioredis";

export interface IUserSessionRepo {
    createSession(idUser: number): Promise<void>
    getUserSessionByToken(token: string): Promise<{
        idUser: number;
        token: string;
    } | null>
}

export class UserSessionRedisRepo implements IUserSessionRepo {
    constructor(
        private readonly connection: Redis
    ) {}

    public async createSession(idUser: number) {
        // Создание токена
        const token = 'asdasdad'
        // Проверка существующего сеанса
        const oldSession = await this.connection.get('user_session:' + idUser)
        if (oldSession) {
            const oldSessionObject: {
                idUser: number,
                token: string
            } = JSON.parse(oldSession)
            // Удаление старого сеанса
            await this.connection.del([ 'user_session:' + idUser ])
            // Удаление обратного ключа старого сеанса
            await this.connection.del([ 'user_session_rev:' + oldSessionObject.token])
        }
        // Создание нового сеанса длительностью в 2 часа
        await this.connection.setex('user_session:' + idUser, 60 * 60 * 2, JSON.stringify({
            idUser, token
        }))
        // Создание обратного ключа по токену длительностью в 2 часа
        await this.connection.setex('user_session_rev:' + token, 60 * 60 * 2, JSON.stringify({
            idUser, token
        }))
    }

    public async getUserSessionByToken(token: string) {
        const session = await this.connection.get('user_session_rev:' + token)
        if (!session) return null
        const sessionObject: {
            idUser: number,
            token: string
        } = JSON.parse(session)
        return sessionObject
    }
}