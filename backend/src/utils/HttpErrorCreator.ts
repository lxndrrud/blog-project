import { Response } from "express";

export interface IHttpErrorCreator {
    badRequest400(res: Response, error: Error): void
    forbidden403(res: Response, error: Error): void
    notFound404(res: Response, error: Error): void
    internalServerError500(res: Response, error: Error): void
}

export class HttpErrorCreator implements IHttpErrorCreator {
    public badRequest400(res: Response, error: Error) {
        res.status(400).json("Неправильный запрос: " + error.message)
    }

    public forbidden403(res: Response, error: Error) {
        res.status(400).json("Запрещено: " + error.message)
    }

    public notFound404(res: Response, error: Error) {
        res.status(404).json("Содержимое не найдено: " + error.message)
    }

    public internalServerError500(res: Response, error: Error) {
        res.status(500).json("Произошла внутренняя ошибка: " + error.message)
    }
}
