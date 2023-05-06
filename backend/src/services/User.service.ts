import { IPermissionRepo } from "../repositories/Permission.repo";
import { IPostRepo } from "../repositories/Post.repo"
import { IUserRepo } from "../repositories/User.repo";
import { IUserSessionRepo } from "../repositories/UserSession.repo";
import { TPermission } from "../types/Permission.type";
import { TPost } from "../types/Post.type";
import { IGenerator } from "../utils/Generator";

export interface IUserService {
    getUserAndPosts(idUser: number): Promise<{
        user: {
            login: string;
        };
        posts: TPost[];
    }>
    loginUser(login: string, password: string): Promise<{
        user: {
            id: number;
            login: string;
        };
        session: {
            idUser: number;
            token: string;
        };
        permissions: TPermission[];
    }>
    registerUser(login: string, password: string): Promise<void>
}

export class UserService implements IUserService {
    constructor(
        private readonly userRepo: IUserRepo,
        private readonly userSessionRepo: IUserSessionRepo,
        private readonly postsRepo: IPostRepo,
        private readonly permissionRepo: IPermissionRepo,
        private readonly generator: IGenerator
    ) {}

    public async getUserAndPosts(idUser: number) {
        const [ user, posts ] = await Promise.all([
            this.userRepo.getUserById(idUser),
            this.postsRepo.getUserPosts(idUser)
        ])
        if (!user) throw new Error('Пользователь не найден...')

        return {
            user: {
                login: user.login
            }, posts
        }
    }

    public async loginUser(login: string, password: string) {
        const user = await this.userRepo.getUserByLogin(login)
        if (!user) throw new Error("Пользователь не найден...")
        const passwordCheck = await this.generator.comparePassword(password, user.password)
        if (!passwordCheck) throw new Error("Пользователь с такими данными для входа не найден!")
        const sessionObject = await this.userSessionRepo.createSession(user.id)
        const permissionList = await this.permissionRepo.getPermissionsList(user.id)
        return {
            user: {
                id: user.id,
                login: user.login
            }, session: sessionObject, permissions: permissionList
        }
    }

    public async registerUser(login: string, password: string) {
        const user = await this.userRepo.getUserByLogin(login)
        if (user) throw new Error("Пользователь с таким логином уже существует!")
        const hashedPassword = await this.generator.hashPassword(password)
        await this.userRepo.insertUser({
            login, password: hashedPassword
        })
    }


}
