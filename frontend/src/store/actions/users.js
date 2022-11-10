import { createAsyncThunk } from "@reduxjs/toolkit"
import axios from "axios"

export const loginRequest = createAsyncThunk(
    "users/login",
    async ({ login, password }, thunkApi) => {
        try {
            const response = await axios.post("/backend/users/login", { login, password })
            return response.data
        } catch (error) {
            throw new Error(error?.response?.data.message || "Произошла непредвиденная ошибка!")
        }
    }
)

export const registerRequest = createAsyncThunk(
    "users/register",
    async ({ login, password }, thunkApi) => {
        try {
            const response = await axios.post("/backend/users/new", { login, password })
            return response.data
        } catch (error) {
            throw new Error(error?.response?.data.message || "Произошла непредвиденная ошибка!")
        }
    }
)
