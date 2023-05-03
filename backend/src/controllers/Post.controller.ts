import { Request, Response } from "express";
import fileUpload from "express-fileupload";


export interface IPostController {

}

export class PostController implements IPostController {
    public async createPost(req: Request, res: Response) {
        const file = <fileUpload.UploadedFile> req.files?.picture
    }
}
