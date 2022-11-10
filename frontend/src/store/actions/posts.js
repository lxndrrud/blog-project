import { createAsyncThunk } from "@reduxjs/toolkit"
import axios from "axios"

const wait = () =>
    new Promise((resolve) => {
        setTimeout(() => resolve(), 500);
});

export const fetchApprovedPosts = createAsyncThunk(
    "posts/fetchApprovedPosts",
    async (_, thunkApi) => {
        try {
            const response = await axios.get("/backend/posts/actual")
            await wait()
            return response.data
        } catch (error) {
            throw new Error(error?.response?.data.message || "Произошла непредвиденная ошибка!")
        }
    }
)

export const createPostRequest = createAsyncThunk(
    'posts/createPost',
    async ({ token, title, annotation, text, timeStart, timeEnd, picture }, thunkApi) => {
        try {
            /*
            let data = new FormData();
            data.append('picture', picture, picture.name);
            data.append('title', title)
            data.append('annotation', annotation)
            data.append('text', text)
            data.append('timeStart', timeStart ? timeStart + " 00:00:00" : null)
            data.append('timeEnd', timeEnd ? timeEnd + " 00:00:00" : null)
            */
            const response = await axios.post("/backend/posts/new", { title, annotation, text, 
                    timeStart: timeStart ? timeStart + " 00:00:00" : null, 
                    timeEnd: timeEnd ? timeEnd + " 00:00:00" : null,
                }, {
                headers: {
                    'auth-token': token,
                    // 'Content-Type': `multipart/form-data`
                }
            })
            return response.data
        } catch (error) {
            throw new Error(error?.response?.data.message || "Произошла непредвиденная ошибка!")
        }
    }
)

export const fetchPostsNeedToApprove = createAsyncThunk(
    'posts/fetchPostsNeedToApprove',
    async ({ token }, thunkApi) => {
        try {
            const response = await axios.get("/backend/posts/needToApprove", {
                headers: {
                    'auth-token': token,
                }
            })
            await wait()
            return response.data
        } catch (error) {
            throw new Error(error?.response?.data.message || "Произошла непредвиденная ошибка!")
        }
    }
)

export const fetchPostNeedToApprove = createAsyncThunk(
    'posts/fetchPostNeedToApprove',
    async ({ token, idPost }, thunkApi) => {
        try {
            const response = await axios.get("/backend/posts/needToApprove/get", {
                params: {
                    idPost 
                },
                headers: {
                    'auth-token': token,
                }
            })
            await wait()
            return response.data
        } catch (error) {
            throw new Error(error?.response?.data.message || "Произошла непредвиденная ошибка!")
        }
    }
)

export const approvePostRequest = createAsyncThunk(
    'posts/approvePost',
    async ({ token, idPost }, thunkApi) => {
        try {
            const response = await axios.post("/backend/posts/approve/", {
                idPost,
            },{
                headers: {
                    'auth-token': token,
                }
            })
            return response.data
        } catch (error) {
            throw new Error(error?.response?.data.message || "Произошла непредвиденная ошибка!")
        }
    }
)

export const rejectPostRequest = createAsyncThunk(
    'posts/rejectPost',
    async ({ token, idPost }, thunkApi) => {
        try {
            const response = await axios.post("/backend/posts/reject/", {
                idPost,
            },{
                headers: {
                    'auth-token': token,
                }
            })
            return response.data
        } catch (error) {
            throw new Error(error?.response?.data.message || "Произошла непредвиденная ошибка!")
        }
    }
)

export const deletePostRequest = createAsyncThunk(
    'posts/deletePost',
    async ({ token, idPost }, thunkApi) => {
        try {
            const response = await axios.delete("/backend/posts/delete/", {
                params: {
                    idPost
                },
                headers: {
                    'auth-token': token,
                }
            })
            return response.data
        } catch (error) {
            throw new Error(error?.response?.data.message || "Произошла непредвиденная ошибка!")
        }
    }
)

export const fetchUserPosts = createAsyncThunk(
    'posts/fetchUserPosts',
    async ({ token }, thunkApi) => {
        try {
            const response = await axios.get("/backend/posts/userPosts/", {
                headers: {
                    'auth-token': token,
                }
            })
            await wait()
            return response.data
        } catch (error) {
            throw new Error(error?.response?.data.message || "Произошла непредвиденная ошибка!")
        }
    }
)

export const fetchApprovedPost = createAsyncThunk(
    'posts/fetchApprovedPost',
    async ({ idPost }, thunkApi) => {
        try {
            const response = await axios.get("/backend/posts/get/", {
                params: {
                    idPost
                }
            })
            return response.data
        } catch (error) {
            throw new Error(error?.response?.data.message || "Произошла непредвиденная ошибка!")
        }
    }
)
