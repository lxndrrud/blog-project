import { createAsyncThunk } from "@reduxjs/toolkit"
import axios from "axios"

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
