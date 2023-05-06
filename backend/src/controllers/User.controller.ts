import { Request, Response } from "express";
import { IUserService } from "../services/User.service";
import { IHttpErrorCreator } from "../utils/HttpErrorCreator";

export interface IUserController {
    getUserAndPosts(req: Request, res: Response): Promise<void>
    loginUser(req: Request, res: Response): Promise<void>
    registerUser(req: Request, res: Response): Promise<void>
}

export class UserController implements IUserController {
    constructor(
        private readonly userService: IUserService,
        private readonly errorCreator: IHttpErrorCreator
    ) {}

    public async getUserAndPosts(req: Request, res: Response) {
        const idUser = <string | undefined> req.query.idUser 
        if (!idUser) {
            return this.errorCreator.badRequest400(res, new Error('Не указан идентификатор пользователя!'))
        }
        const pIdUser = parseInt(idUser)
        try {
            const info = await this.userService.getUserAndPosts(pIdUser)
            res.status(200).json(info)
        } catch(error: any) {
            if (error instanceof Error) {
                this.errorCreator.internalServerError500(res, error)
            }
        }
    }

    public async loginUser(req: Request, res: Response) {
        let payload: {
            login: string,
            password: string
        }
        try {
            payload = {
                login: req.body.login,
                password: req.body.password
            }
        } catch(error: any) {
            return this.errorCreator.badRequest400(res, error)
        }
        try {
            const info = await this.userService.loginUser(payload.login, payload.password)
            res.status(200).json({
                token: info.session.token,
                currentUser: info.user,
                permissions: info.permissions
            })
        } catch(error: any) {
            if (error instanceof Error) {
                this.errorCreator.internalServerError500(res, error)
            }
        }
    }

    public async registerUser(req: Request, res: Response) {
        let payload: {
            login: string,
            password: string
        }
        try {
            payload = {
                login: req.body.login,
                password: req.body.password
            }
        } catch(error: any) {
            return this.errorCreator.badRequest400(res, error)
        }
        try {
            await this.userService.registerUser(payload.login, payload.password)
            res.status(201).end()
        } catch(error: any) {
            if (error instanceof Error) {
                this.errorCreator.internalServerError500(res, error)
            }
        }
    }

}
