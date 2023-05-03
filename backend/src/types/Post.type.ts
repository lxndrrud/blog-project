
export type TPost = {
    id: number;
    id_author: number;
    author_login: string;
    title: string;
    annotation: string;
    text: string;
    picture: string;
    views: number;
    approved: boolean;
    rejected: boolean;
    time_start: Date;
    time_end: Date;
    created_at: Date;
}
