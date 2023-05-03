import fileUpload from 'express-fileupload'
import fs from 'fs/promises'
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
        const filename = await this.generator.generateRandomString(15)
        await file.mv(`../../storage/${filename}`)
        return filename
    }

    public async deleteFile(filename: string) {
        await fs.unlink(`../../storage/${filename}`)
    }
}
