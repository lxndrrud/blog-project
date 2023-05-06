import { Router } from "express";
import { PostRouter } from "./Post.router";
import { UserRouter } from "./User.router";


export const MainRouter = Router()
const prefixRouter = Router()

MainRouter.use('/backend', prefixRouter)

prefixRouter.use('/users', UserRouter)
prefixRouter.use('/posts', PostRouter)
