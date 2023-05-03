import bcrypt from 'bcrypt'

export interface IGenerator {
    generateToken(idUser: number): Promise<string>
    hashPassword(password: string): Promise<string>
    comparePassword(password: string, passwordHash: string): Promise<boolean>
    generateRandomString(length: number): Promise<string>
}

export class Generator implements IGenerator {
    public async generateToken(idUser: number) {
        return bcrypt.hash(idUser.toString(), 10)
    }

    public async hashPassword(password: string) {
        return bcrypt.hash(password, 10)
    }

    public async comparePassword(password: string, passwordHash: string) {
        return bcrypt.compare(password, passwordHash)
    }

    public async generateRandomString(length: number) {
        let result = '';
        const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
        const charactersLength = characters.length;
        let counter = 0;
        while (counter < length) {
            result += characters.charAt(Math.floor(Math.random() * charactersLength));
            counter += 1;
        }
        return result;
    }
}


