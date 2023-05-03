import { IPermissionRepo } from "../repositories/Permission.repo";
import { IUserSessionRepo } from "../repositories/UserSession.repo";
import { TPermission } from "../types/Permission.type";

export interface IPermissionInfra {
    getPermissionList(token: string): Promise<TPermission[]>
}

export class PermissionInfra implements IPermissionInfra {
    constructor(
        private readonly userSessionRepo: IUserSessionRepo,
        private readonly permissionRepo: IPermissionRepo
    ) {}

    public async getPermissionList(token: string) {
        const session = await this.userSessionRepo.getUserSessionByToken(token)
        if (!session) throw new Error("Сеанс не найден. Пожалуйста, войдите по новой.")
        const permissionList = await this.permissionRepo.getPermissionsList(session.idUser)
        return permissionList
    }
}
