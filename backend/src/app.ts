import bodyParser from "body-parser"
import fileUpload from 'express-fileupload'
import express from "express"
import { MainRouter } from "./routes/Main.router"


export async function Bootstrap() {
    const app = express()

    app.use(bodyParser.json())

    app.use(fileUpload())

    app.use(MainRouter)

    app.use((req, res) => {
        res.status(404).json({
            message: 'Страница не найдена!'
        })
    })

    return app
}