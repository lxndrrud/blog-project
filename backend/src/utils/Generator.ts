import bcrypt from 'bcrypt'

export interface IGenerator {
    generateToken(idUser: number): Promise<string>
    hashPassword(password: string): Promise<string>
    comparePassword(password: string, passwordHash: string): Promise<boolean>
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
}


