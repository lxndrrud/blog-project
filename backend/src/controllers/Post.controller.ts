import { Request, Response } from "express";
import fileUpload from "express-fileupload";
import { IPostService } from "../services/Post.service";
import { IHttpErrorCreator } from "../utils/HttpErrorCreator";

export interface IPostController {
    getUserPosts(req: Request, res: Response): Promise<void>
    getApprovedPosts(req: Request, res: Response): Promise<void>
    getApprovedPost(req: Request, res: Response): Promise<void>
    getPostsNeedToApprove(req: Request, res: Response): Promise<void>
    getPostNeedToApprove(req: Request, res: Response): Promise<void>
    createPost(req: Request, res: Response): Promise<void>
    approvePost(req: Request, res: Response): Promise<void>
    rejectPost(req: Request, res: Response): Promise<void>
    deletePost(req: Request, res: Response): Promise<void>
}

export class PostController implements IPostController {
    constructor(
        private readonly postService: IPostService,
        private readonly errorCreator: IHttpErrorCreator
    ) {}

    public async getUserPosts(req: Request, res: Response) {
        const token = req.header("auth-token")
        if (!token) {
            return this.errorCreator.forbidden403(res, new Error('Не указан токен доступа!'))
        }
        try {
            const posts = await this.postService.getUserPosts(token)
            res.status(200).json({
                posts
            })
        } catch(error: any) {
            if (error instanceof Error) {
                this.errorCreator.internalServerError500(res, error)
            }
        }
    }

    public async getApprovedPosts(req: Request, res: Response) {
        try {
            const posts = await this.postService.getApprovedPosts()
            res.status(200).json({
                posts
            })
        } catch(error: any) {
            if (error instanceof Error) {
                this.errorCreator.internalServerError500(res, error)
            }
        }
    }

    public async getApprovedPost(req: Request, res: Response) {
        const idPost = <string | undefined> req.query.idPost 
        if (!idPost) {
            return this.errorCreator.badRequest400(res, new Error('Не указан идентификатор поста!'))
        }
        const pIdPost = parseInt(idPost)
        try {
            const post = await this.postService.getApprovedPost(pIdPost)
            res.status(200).json({
                post
            })
        } catch(error: any) {
            if (error instanceof Error) {
                this.errorCreator.internalServerError500(res, error)
            }
        }
    }

    public async createPost(req: Request, res: Response) {
        const token = req.header('auth-token')
        if (!token) {
            return this.errorCreator.forbidden403(res, new Error('Не указан токен доступа!'))
        }
        let requestBody: {
            title: string,
            text: string,
            annotation: string,
            time_start?: Date,
            time_end?: Date,
            picture: fileUpload.UploadedFile
        }
        try {
            requestBody = {
                title: req.body.title,
                text: req.body.text,
                annotation: req.body.annotation,
                time_start: req.body.timeStart ? new Date(<string> req.body.timeStart) : undefined,
                time_end: req.body.timeEnd ? new Date(<string> req.body.timeEnd) : undefined,
                picture: <fileUpload.UploadedFile>req.files?.picture 
            }
        } catch(error: any) {
            return this.errorCreator.badRequest400(res, error)
        }
        try {
            await this.postService.insertPost(token, requestBody)
            res.status(201).end()
        } catch(error: any) {
            if (error instanceof Error) {
                this.errorCreator.internalServerError500(res, error)
            }
        }
    }

    public async getPostsNeedToApprove(req: Request, res: Response) {
        const token = req.header("auth-token")
        if (!token) {
            return this.errorCreator.forbidden403(res, new Error('Не указан токен доступа!'))
        }
        try {
            const posts = await this.postService.getPostsNeedToApprove(token)
            res.status(200).json({
                posts
            })
        } catch(error: any) {
            if (error instanceof Error) {
                this.errorCreator.internalServerError500(res, error)
            }
        }
    }

    public async getPostNeedToApprove(req: Request, res: Response) {
        const token = req.header("auth-token")
        if (!token) {
            return this.errorCreator.forbidden403(res, new Error('Не указан токен доступа!'))
        }
        const idPost = <string | undefined> req.query.idPost 
        if (!idPost) {
            return this.errorCreator.badRequest400(res, new Error('Не указан идентификатор поста!'))
        }
        const pIdPost = parseInt(idPost)
        try {
            const post = await this.postService.getPostNeedToApprove(token, pIdPost)
            res.status(200).json({
                post
            })
        } catch(error: any) {
            if (error instanceof Error) {
                this.errorCreator.internalServerError500(res, error)
            }
        }
    }

    public async approvePost(req: Request, res: Response) {
        const token = req.header("auth-token")
        if (!token) {
            return this.errorCreator.forbidden403(res, new Error('Не указан токен доступа!'))
        }
        const idPost: number | undefined = req.body.idPost 
        if (!idPost) {
            return this.errorCreator.badRequest400(res, new Error('Не указан идентификатор поста!'))
        }
        try {
            await this.postService.approvePost(token, idPost)
            res.status(200).end()
        } catch(error: any) {
            if (error instanceof Error) {
                this.errorCreator.internalServerError500(res, error)
            }
        }
    }

    public async rejectPost(req: Request, res: Response) {
        const token = req.header("auth-token")
        if (!token) {
            return this.errorCreator.forbidden403(res, new Error('Не указан токен доступа!'))
        }
        const idPost: number | undefined = req.body.idPost 
        if (!idPost) {
            return this.errorCreator.badRequest400(res, new Error('Не указан идентификатор поста!'))
        }
        try {
            await this.postService.rejectPost(token, idPost)
            res.status(200).end()
        } catch(error: any) {
            if (error instanceof Error) {
                this.errorCreator.internalServerError500(res, error)
            }
        }
    }

    public async deletePost(req: Request, res: Response) {
        const token = req.header("auth-token")
        if (!token) {
            return this.errorCreator.forbidden403(res, new Error('Не указан токен доступа!'))
        }
        const idPost = <string | undefined> req.query.idPost 
        if (!idPost) {
            return this.errorCreator.badRequest400(res, new Error('Не указан идентификатор поста!'))
        }
        const pIdPost = parseInt(idPost)
        try {
            await this.postService.deletePost(token, pIdPost)
            res.status(200).end()
        } catch(error: any) {
            if (error instanceof Error) {
                this.errorCreator.internalServerError500(res, error)
            }
        }
    }
}
