import { Router } from "express";
import { postController } from "../container";

export const PostRouter = Router()

PostRouter.get('/userPosts', postController.getUserPosts.bind(postController))
PostRouter.get('/actual', postController.getApprovedPosts.bind(postController))
PostRouter.get('/needToApprove', postController.getPostsNeedToApprove.bind(postController))
PostRouter.get('/get', postController.getApprovedPost.bind(postController))
PostRouter.get('/needToApprove/get', postController.getApprovedPost.bind(postController))
PostRouter.post('/new', postController.createPost.bind(postController))
PostRouter.post('/approve', postController.approvePost.bind(postController))
PostRouter.post('/reject', postController.rejectPost.bind(postController))
PostRouter.delete('/delete', postController.deletePost.bind(postController))
