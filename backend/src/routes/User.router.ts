import { Router } from "express";
import { userController } from "../container";

export const UserRouter = Router()

UserRouter.get('/profile', userController.getUserAndPosts.bind(userController))
UserRouter.post('/login', userController.loginUser.bind(userController))
UserRouter.post('/new', userController.registerUser.bind(userController))
