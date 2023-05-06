import fileUpload from "express-fileupload";
import { IPermissionInfra } from "../infrastructure/Permission.infra";
import { IPostRepo } from "../repositories/Post.repo";
import { IUserSessionRepo } from "../repositories/UserSession.repo";
import { TPost } from "../types/Post.type";
import { IFileProcessor } from "../utils/FileProcessor";
import { IPermissionChecker } from "../utils/PermissionChecker";

export interface IPostService {
    getUserPosts(token: string): Promise<TPost[]>
    getApprovedPosts(): Promise<TPost[]>
    getApprovedPost(idPost: number): Promise<TPost>
    getPostsNeedToApprove(token: string): Promise<TPost[]>
    getPostNeedToApprove(token: string, idPost: number): Promise<TPost>
    insertPost(token: string, payload: {
        title: string;
        text: string;
        annotation: string;
        picture: fileUpload.UploadedFile;
        time_start?: Date;
        time_end?: Date;
    }): Promise<void>
    approvePost(token: string, idPost: number): Promise<void>
    rejectPost(token: string, idPost: number): Promise<void>
    deletePost(token: string, idPost: number): Promise<void>
}

export class PostService implements IPostService {
    constructor(
        private readonly permissionInfra: IPermissionInfra,
        private readonly permissionChecker: IPermissionChecker,
        private readonly postsRepo: IPostRepo,
        private readonly userSessionRepo: IUserSessionRepo,
        private readonly fileProcessor: IFileProcessor
    ) {}

    public async getUserPosts(token: string) {
        const sessionObject = await this.userSessionRepo.getUserSessionByToken(token)
        if (!sessionObject) throw new Error("Сеанс не найден. Пожалуйста, войдите по новой.")
        const posts = await this.postsRepo.getUserPosts(sessionObject.idUser, true)
        return posts
    }

    public async getApprovedPosts() {
        const posts = await this.postsRepo.getApprovedPosts()
        return posts
    }

    public async getApprovedPost(idPost: number) {
        const post = await this.postsRepo.getApprovedPost(idPost)
        if (!post) throw new Error('Пост не найден...')
        await this.postsRepo.addViews(idPost, 1)
        post.views += 1
        return post
    }

    public async getPostsNeedToApprove(token: string) {
        const permissions = await this.permissionInfra.getPermissionList(token)
        if (!this.permissionChecker.canModeratePosts(permissions)) 
            throw new Error("Вам запрещено модерировать посты!")
        const posts = await this.postsRepo.getPostsNeedToApprove()
        return posts
    }

    public async getPostNeedToApprove(token: string, idPost: number) {
        const permissions = await this.permissionInfra.getPermissionList(token)
        if (!this.permissionChecker.canModeratePosts(permissions)) 
            throw new Error("Вам запрещено модерировать посты!")
        const post = await this.postsRepo.getPostNeedToApprove(idPost)
        if (!post) throw new Error("Пост не найден...")
        return post
    }

    public async insertPost(token: string, payload: {
        title: string;
        text: string;
        annotation: string;
        picture: fileUpload.UploadedFile;
        time_start?: Date;
        time_end?: Date;
    }) {
        const permissions = await this.permissionInfra.getPermissionList(token)
        if (!this.permissionChecker.canCreatePosts(permissions))
            throw new Error("Вам запрещено создавать посты!")
        const sessionObject = await this.userSessionRepo.getUserSessionByToken(token)
        if (!sessionObject) throw new Error("Сеанс не найден. Пожалуйста, войдите по новой.")
        const filename = await this.fileProcessor.saveFileOnDisk(payload.picture)
        await this.postsRepo.insertPost({
            ...payload,
            id_author: sessionObject.idUser,
            picture: filename
        })
    }

    public async approvePost(token: string, idPost: number) {
        const permissions = await this.permissionInfra.getPermissionList(token)
        if (!this.permissionChecker.canModeratePosts(permissions)) 
            throw new Error("Вам запрещено модерировать посты!")
        await this.postsRepo.approvePost(idPost)
    }

    public async rejectPost(token: string, idPost: number) {
        const permissions = await this.permissionInfra.getPermissionList(token)
        if (!this.permissionChecker.canModeratePosts(permissions)) 
            throw new Error("Вам запрещено модерировать посты!")
        await this.postsRepo.rejectPost(idPost)
    }

    public async deletePost(token: string, idPost: number) {
        const sessionObject = await this.userSessionRepo.getUserSessionByToken(token)
        if (!sessionObject) throw new Error("Сеанс не найден. Пожалуйста, войдите по новой.")
        const post = await this.postsRepo.getPostById(idPost)
        if (!post) throw new Error("Пост не найден...")
        if (post.id_author !== sessionObject.idUser) 
            throw new Error("Вы не являетесь автором поста!")
        await this.fileProcessor.deleteFile(post.picture)
        await this.postsRepo.deletePost(idPost)
    }
}
