import { Knex } from "knex";

export interface IPermissionRepo {
    getPermissionsList(idUser: number): Promise<{
        id: number;
        title: string;
        code: string;
    }[]>
}

export class PermissionPostgresRepo implements IPermissionRepo {
    constructor(
        private readonly connection: Knex
    ) {}

    public async getPermissionsList(idUser: number) {
        const permissions: {
            id: number,
            title: string,
            code: string
        }[] = await this.connection('public.users as u')
            .innerJoin('public.roles as r', 'r.id', '=', 'u.id_role')
            .innerJoin('public.roles_permissions as rp', 'rp.id_role', '=', 'r.id')
            .innerJoin('public.permissions as p', 'p.id', '=', 'rp.id_permission')
            .select('p.*')
            .where('u.id', idUser)
        return permissions
    }
}
