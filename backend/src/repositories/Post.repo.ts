import { Knex } from "knex";
import { TPost } from "../types/Post.type";

export interface IPostRepo {
    getPostById(idPost: number): Promise<TPost | undefined>
    getApprovedPosts(): Promise<TPost[]>
    getApprovedPost(idPost: number): Promise<TPost | undefined>
    getPostsNeedToApprove(): Promise<TPost[]>
    getPostNeedToApprove(idPost: number): Promise<TPost | undefined>
    getUserPosts(idUser: number, isOwner?: boolean): Promise<TPost[]>
    insertPost(payload: {
        title: string;
        text: string;
        annotation: string;
        picture: string;
        id_author: number;
        time_start?: Date;
        time_end?: Date;
    }): Promise<void>
    addViews(idPost: number, viewsQuantity: number): Promise<void>
    approvePost(idPost: number): Promise<void>
    rejectPost(idPost: number): Promise<void>
    deletePost(idPost: number): Promise<void>
}

export class PostPostgresRepo implements IPostRepo {
    constructor(
        private readonly connection: Knex
    ) {}

    public async getPostById(idPost: number) {
        const post: TPost | undefined = await this.connection('public.posts as p')
            .select('p.*', 'u.login as author_login')
            .innerJoin('public.users as u', 'u.id', 'p.id_author')
            .where('p.id', idPost)
            .first()
        return post
    }

    public async getApprovedPosts() {
        const posts: TPost[] = await this.connection('public.posts as p')
            .select('p.*', 'u.login as author_login')
            .innerJoin('public.users as u', 'u.id', 'p.id_author')
            .where('p.approved', true)
            .andWhere(qb => {
                qb.where('p.time_start IS NULL')
                qb.orWhere('p.time_start <= NOW()')
            })
            .andWhere(qb => {
                qb.where('p.time_end IS NULL')
                qb.orWhere('p.time_end > NOW()')
            })
            .orderBy('p.views', 'desc')
        return posts
    }

    public async getApprovedPost(idPost: number) {
        const post: TPost | undefined = await this.connection('public.posts as p')
            .select('p.*', 'u.login as author_login')
            .innerJoin('public.users as u', 'u.id', 'p.id_author')
            .where('p.id', idPost)
            .andWhere('p.approved', true)
            .andWhere(qb => {
                qb.where('p.time_start IS NULL')
                qb.orWhere('p.time_start <= NOW()')
            })
            .andWhere(qb => {
                qb.where('p.time_end IS NULL')
                qb.orWhere('p.time_end > NOW()')
            })
            .orderBy('p.views', 'desc')
            .first()
        return post
    }

    public async getPostsNeedToApprove() {
        const posts: TPost[] = await this.connection('public.posts as p')
            .select('p.*', 'u.login as author_login')
            .innerJoin('public.users as u', 'u.id', 'p.id_author')
            .where('p.approved', false)
            .andWhere('p.rejected', false)
            .andWhere(qb => {
                qb.where('p.time_start IS NULL')
                qb.orWhere('p.time_start <= NOW()')
            })
            .andWhere(qb => {
                qb.where('p.time_end IS NULL')
                qb.orWhere('p.time_end > NOW()')
            })
        return posts
    }

    public async getPostNeedToApprove(idPost: number) {
        const post: TPost | undefined = await this.connection('public.posts as p')
            .select('p.*', 'u.login as author_login')
            .innerJoin('public.users as u', 'u.id', 'p.id_author')
            .where('p.id', idPost)
            .andWhere('p.approved', false)
            .andWhere('p.rejected', false)
            .first()
        return post
    }

    public async getUserPosts(idUser: number, isOwner?: boolean) {
        const posts: TPost[] = await this.connection('public.posts as p')
            .select('p.*', 'u.login as author_login')
            .innerJoin('public.users as u', 'u.id', 'p.id_author')
            .where('u.id', idUser)
            .andWhere(qb => {
                if (!isOwner) {
                    qb.where('p.approved', true)
                    qb.andWhere(qb1 => {
                        qb1.where('p.time_start IS NULL')
                        qb1.orWhere('p.time_start <= NOW()')
                    })
                    qb.andWhere(qb1 => {
                        qb1.where('p.time_end IS NULL')
                        qb1.orWhere('p.time_end > NOW()')
                    })
                }
            })
            .orderBy([
                { column: 'p.created_at', order: 'desc' },
                { column: 'p.views', order: 'desc' },
            ])
        return posts
    }

    public async insertPost(payload: { 
        title: string,
        text: string,
        annotation: string,
        picture: string,
        id_author: number,
        time_start?: Date,
        time_end?: Date
    }) {
        await this.connection('public.posts')
            .insert(payload)
    }

    public async addViews(idPost: number, viewsQuantity: number) {
        await this.connection('public.posts as p')
            .where('p.id', idPost)
            .update({
                views: "views + " + viewsQuantity.toString()
            })
    }

    public async approvePost(idPost: number) {
        await this.connection('public.posts as p')
            .where('p.id', idPost)
            .andWhere('p.rejected', false)
            .update({
                approved: true
            })
    }

    public async rejectPost(idPost: number) {
        await this.connection('public.posts as p')
            .where('p.id', idPost)
            .andWhere('p.approved', false)
            .update({
                rejected: true
            })
    }

    public async deletePost(idPost: number) {
        await this.connection('public.posts as p')
            .where('p.id', idPost)
            .delete()
    }
}
