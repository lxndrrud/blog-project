import fileUpload from 'express-fileupload'
import fs from 'fs/promises'
import path from 'path'
import { IGenerator } from './Generator'

export interface IFileProcessor {
    saveFileOnDisk(file: fileUpload.UploadedFile): Promise<string>
    deleteFile(filename: string): Promise<void>
}

export class FileProcessor implements IFileProcessor {
    constructor(
        private readonly generator: IGenerator
    ) {}

    public async saveFileOnDisk(file: fileUpload.UploadedFile) {
        const fileExtension = file.name.split('.')[1]
        const filename = await this.generator.generateRandomString(15)
        await file.mv(path.join(__dirname, '..', '..', 'storage', `${filename}.${fileExtension}`))
        return `/${filename}.${fileExtension}`
    }

    public async deleteFile(filename: string) {
        // Название файла содержит начальный слэш
        await fs.unlink(path.join(__dirname, '..', '..', 'storage', filename.slice(1)))
    }
}
