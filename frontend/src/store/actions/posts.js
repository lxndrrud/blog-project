import { createAsyncThunk } from "@reduxjs/toolkit"
import axios from "axios"
import moment from 'moment'

export const fetchApprovedPosts = createAsyncThunk(
    "posts/fetchApprovedPosts",
    async (_, thunkApi) => {
        try {
            const response = await axios.get("/backend/posts/actual")
            return response.data
        } catch (error) {
            throw new Error(error?.response?.data.message || "Произошла непредвиденная ошибка!")
        }
    }
)

export const createPostRequest = createAsyncThunk(
    'posts/createPost',
    async ({ token, title, annotation, text, timeStart, timeEnd }, thunkApi) => {
        try {
            const response = await axios.post("/backend/posts/new", { title, annotation, text, 
                    timeStart: timeStart ? timeStart + " 00:00:00" : null, 
                    timeEnd: timeEnd ? timeEnd + " 00:00:00" : null,
                }, {
                headers: {
                    'auth-token': token
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
