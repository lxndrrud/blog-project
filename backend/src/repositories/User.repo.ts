import { Knex } from "knex";

export interface IUserRepo {
    getUserById(idUser: number): Promise<{
        id: number;
        login: string;
        password: string;
    } | undefined>
    getUserByLogin(login: string): Promise<{
        id: number;
        login: string;
        password: string;
    } | undefined>
    insertUser(payload: {
        login: string;
        password: string;
    }): Promise<number[]>
}

export class UserPostgresRepo implements IUserRepo {
    constructor(
        private readonly connection: Knex
    ) {}

    public async getUserById(idUser: number) {
        const user: {
            id: number,
            login: string,
            password: string
        } | undefined = await this.connection('public.users as u')
            .where('u.id', idUser)
            .first()
        return user
    }

    public async getUserByLogin(login: string) {
        const user: {
            id: number,
            login: string,
            password: string
        } | undefined = await this.connection('public.users as u')
            .where('u.login', login)
            .first()
        return user
    }

    public async insertUser(payload: { login: string, password: string }) {
        return this.connection('public.users')
            .returning('id')
            .insert({
                ...payload, id_role: 2
            })
    }
}
